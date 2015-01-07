package tpl

import (
	"fmt"
	"text/template"
)

type tpl struct {
	Name     string
	Template string
}

type tpls map[string]Tpl

var templates tpls = tpls{
	"module.info": tpl{
		"module.info",
		`
name = {{.Name}}
description = {{.Desc}}
package = {{.Package}}

core = {{.Core}}
version = {{.Version}}
		`,
	},
}

func init() {

}
