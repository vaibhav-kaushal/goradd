package generator

import (
	"github.com/goradd/goradd/pkg/config"
	generator3 "github.com/goradd/goradd/pkg/page/control/generator"
)

func init() {
	if !config.Release {
		//generator2.RegisterControlGenerator(Checkbox{})
	}
}

// This structure describes the textbox to the connector dialog and code generator
type Checkbox struct {
	generator3.Checkbox // base it on the built-in generator
}

func (d Checkbox) Imports() []string {
	return []string{"github.com/goradd/goradd/pkg/bootstrap/control"}
}
