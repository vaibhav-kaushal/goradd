package generator

import (
	generator2 "github.com/goradd/goradd/codegen/generator"
	"github.com/goradd/goradd/pkg/config"
	generator3 "github.com/goradd/goradd/pkg/page/control/generator"
)

func init() {
	if !config.Release {
		generator2.RegisterControlGenerator(IntegerTextbox{})
	}
}

// This structure describes the textbox to the connector dialog and code generator
type IntegerTextbox struct {
	generator3.IntegerTextbox // base it on the built-in generator
}

func (d IntegerTextbox) Imports() []string {
	return []string{"github.com/goradd/goradd/pkg/bootstrap/control"}
}
