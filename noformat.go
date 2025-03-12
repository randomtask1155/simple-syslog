package main

import (
	"bufio"
	"time"

	"gopkg.in/mcuadros/go-syslog.v2/format"
)

var noFormat = &NoFormat{}

type NoFormat struct{}

func (f NoFormat) GetParser(line []byte) format.LogParser {
	return &noFormatParser{string(line)}
}

func (f NoFormat) GetSplitFunc() bufio.SplitFunc {
	return nil // not used
}

type noFormatParser struct {
	line string
}

func (c noFormatParser) Dump() format.LogParts {
	return format.LogParts{
		"content": string(c.line),
	}
}

func (c noFormatParser) Parse() error {
	return nil // doesn't parse anything
}

func (c noFormatParser) Location(location *time.Location) {
	// not used
}
