package main

import (
	"github.com/jwaldrip/odin/cli"
)

// CLI is the odin CLI
var app = cli.New("0.0.1", "Elevated Third CLI Utility", func(c cli.Command) {
})

var newCmd = commandNew()

//func init() {
//}

func main() {
	app.Start()
}
