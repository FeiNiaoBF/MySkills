package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"researchcurator/curator"
	"researchcurator/publisher"
	"researchcurator/visualizer"
)

func main() {
	if e := execute(os.Args[1:], os.Stdin, os.Stdout, os.Stderr); e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
}
func execute(args []string, in io.Reader, out, errs io.Writer) error {
	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		fmt.Fprintln(out, "researchcurator <validate|validate-report|import|dedup|rank|finalize|render|record|record-round|publish> [-in input.json|-] [-out output|-]\nrecord requires -kind event|query|source|decision|adjudication -record record.json; query/source capture timestamps when recorded. record-round requires -record round.json and resets the previous stop. finalize writes the legacy report. validate-report checks a phased handoff; default publish path is <system-temp>/research-curator/<topic>/index.html. No network access; retrieval uses host tools.")
		return nil
	}
	command := args[0]
	f := flag.NewFlagSet(command, flag.ContinueOnError)
	f.SetOutput(errs)
	input := f.String("in", "-", "run JSON file, or - for stdin")
	output := f.String("out", "-", "output file, or - for stdout")
	report := f.String("report", "", "finalize HTML report; default report.html next to file output")
	kind := f.String("kind", "", "record kind")
	record := f.String("record", "", "record JSON file")
	if e := f.Parse(args[1:]); e != nil {
		return e
	}
	if f.NArg() != 0 {
		return fmt.Errorf("unexpected positional arguments")
	}
	if *report != "" && command != "finalize" {
		return fmt.Errorf("-report requires finalize")
	}
	if command == "finalize" {
		if *report == "" {
			if *output == "-" {
				return fmt.Errorf("finalize requires -out file or explicit -report file for stdout")
			}
			*report = filepath.Join(filepath.Dir(*output), "report.html")
		}
		if *report == "-" {
			return fmt.Errorf("report must be a file")
		}
		if *output != "-" {
			a, _ := filepath.Abs(*output)
			b, _ := filepath.Abs(*report)
			if a == b {
				return fmt.Errorf("JSON and HTML outputs must differ")
			}
		}
	}
	var reader io.Reader = in
	var inputFile *os.File
	if *input != "-" {
		file, err := os.Open(*input)
		if err != nil {
			return err
		}
		inputFile = file
		reader = inputFile
	}
	b, e := io.ReadAll(io.LimitReader(reader, 32*1024*1024+1))
	if inputFile != nil {
		if closeErr := inputFile.Close(); e == nil {
			e = closeErr
		}
	}
	if e != nil {
		return e
	}
	if len(b) > 32*1024*1024 {
		return fmt.Errorf("input exceeds 32 MiB")
	}
	if command == "record-round" {
		if *record == "" {
			return fmt.Errorf("record-round requires -record round.json")
		}
		report, err := curator.DecodeReport(b)
		if err != nil {
			return err
		}
		file, err := os.Open(*record)
		if err != nil {
			return err
		}
		defer file.Close()
		var round curator.ResearchRound
		d := json.NewDecoder(io.LimitReader(file, 1024*1024))
		d.DisallowUnknownFields()
		if err := d.Decode(&round); err != nil {
			return err
		}
		if err := d.Decode(new(any)); err != io.EOF {
			return fmt.Errorf("trailing round JSON")
		}
		if err := curator.RecordResearchRound(report, round); err != nil {
			return err
		}
		data, err := json.MarshalIndent(report, "", "  ")
		if err != nil {
			return err
		}
		data = append(data, '\n')
		if *output != "-" {
			return curator.WriteFileAtomic(*output, data)
		}
		_, err = out.Write(data)
		return err
	}
	if command == "validate-report" {
		if *output != "-" {
			return fmt.Errorf("validate-report does not accept -out")
		}
		report, err := curator.DecodeReport(b)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(out, "valid report envelope %s\n", report.ReportVersion)
		return err
	}
	if command == "publish" {
		destination := *output
		defaultDestination := destination == "-"
		if defaultDestination {
			report, err := curator.DecodeReport(b)
			if err != nil {
				return err
			}
			topic := report.Run.Contract.Question
			if report.Article != nil && report.Article.Title != "" {
				topic = report.Article.Title
			}
			destination, err = publisher.DefaultDestination(topic, os.TempDir())
			if err != nil {
				return err
			}
		}
		if err := publisher.Write(b, destination); err != nil {
			return err
		}
		if defaultDestination {
			_, err := fmt.Fprintln(out, destination)
			return err
		}
		return nil
	}
	r, e := curator.Decode(b)
	if e != nil {
		return e
	}
	var data []byte
	switch command {
	case "validate", "import":
	case "dedup":
		if e := curator.Deduplicate(r); e != nil {
			return e
		}
	case "rank":
		selected := curator.Rank(r)
		for _, s := range r.Sources {
			if s.Status != "selected" {
				selected = append(selected, s)
			}
		}
		r.Sources = selected
	case "finalize":
		if e := curator.Finalize(r); e != nil {
			return e
		}
	case "render":
		data, e = visualizer.Render(b)
		if e != nil {
			return e
		}
	case "record":
		if *record == "" {
			return fmt.Errorf("record requires -record file")
		}
		file, e := os.Open(*record)
		if e != nil {
			return e
		}
		defer file.Close()
		d := json.NewDecoder(io.LimitReader(file, 1024*1024))
		d.DisallowUnknownFields()
		rec := curator.Recorder{Run: r}
		switch *kind {
		case "event":
			var x curator.Event
			if e := d.Decode(&x); e != nil {
				return e
			}
			e = rec.RecordEvent(x)
		case "query":
			var x curator.Query
			if e := d.Decode(&x); e != nil {
				return e
			}
			e = rec.RecordQueryAtCapture(x)
		case "source":
			var x curator.Source
			if e := d.Decode(&x); e != nil {
				return e
			}
			e = rec.RecordSourceAtCapture(x)
		case "decision":
			var x curator.Decision
			if e := d.Decode(&x); e != nil {
				return e
			}
			e = rec.RecordDecision(x)
		case "adjudication":
			var x curator.ConflictAdjudication
			if e := d.Decode(&x); e != nil {
				return e
			}
			e = rec.RecordAdjudication(x)
		default:
			return fmt.Errorf("record requires -kind event|query|source|decision|adjudication")
		}
		if e != nil {
			return e
		}
		if e := d.Decode(new(any)); e != io.EOF {
			return fmt.Errorf("trailing record JSON")
		}
	default:
		return fmt.Errorf("unknown command %q", command)
	}
	if command != "render" {
		data, e = json.MarshalIndent(r, "", "  ")
		if e != nil {
			return e
		}
		data = append(data, '\n')
	}
	if command == "finalize" {
		html, err := visualizer.Render(data)
		if err != nil {
			return err
		}
		if err := curator.WriteFileAtomic(*report, html); err != nil {
			return err
		}
	}
	if *output != "-" {
		return curator.WriteFileAtomic(*output, data)
	}
	_, e = out.Write(data)
	return e
}
