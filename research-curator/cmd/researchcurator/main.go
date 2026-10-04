package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"researchcurator/curator"
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
		fmt.Fprintln(out, "researchcurator <validate|import|dedup|rank|finalize|render|record> [-in run.json|-] [-out output|-]\nrecord also requires -kind event|query|decision -record record.json. finalize writes an automatic report: -out run.json defaults to adjacent report.html; stdout requires -report report.html. No network access; research is supplied by retrieval tools.")
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
	if *input != "-" {
		file, e := os.Open(*input)
		if e != nil {
			return e
		}
		defer file.Close()
		reader = file
	}
	b, e := io.ReadAll(io.LimitReader(reader, 32*1024*1024+1))
	if e != nil {
		return e
	}
	if len(b) > 32*1024*1024 {
		return fmt.Errorf("input exceeds 32 MiB")
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
			e = rec.RecordQuery(x)
		case "decision":
			var x curator.Decision
			if e := d.Decode(&x); e != nil {
				return e
			}
			e = rec.RecordDecision(x)
		default:
			return fmt.Errorf("record requires -kind event|query|decision")
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
