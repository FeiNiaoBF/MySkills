package schema

import _ "embed"

//go:embed run.schema.json
var Run []byte

//go:embed report.schema.json
var Report []byte
