package main

import (
	"os"
	"fmt"
	"github.com/jwaldrip/odin/cli"
)

func commandNew() *cli.SubCommand {
	var cmd = app.DefineSubCommand(
		"new",
		"Utilities for working with modules.",
		func(c cli.Command) {},
	)

	module := cmd.DefineSubCommand("module", "Creates a new module.", newModule, "name")
	module.DefineStringFlag(
		"directory",
		"sites/all/modules/custom",
		"A directory path relative to the current Drupal root. Defaults to sites/all/modules/custom.",
	)

	cmd.DefineSubCommand("site", "Creates a new site skel.", newSite, "name")

	return cmd
}

func newModule(c cli.Command) {
	var name = c.Param("name")

	root, err := drupalRoot()

	if err != nil {
		fmt.Println(err)
		fmt.Println("You may want to set the --directory flag.")
		return
	}

	dir := fmt.Sprintf("%s/%s/%s", root, c.Flag("directory"), name)

	err = os.MkdirAll(dir, 0755)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func newSite(c cli.Command) {}
