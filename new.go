package main

import (
	"fmt"
	"github.com/jwaldrip/odin/cli"
	"os"
	"path"
	"strings"
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
		"A directory path relative to the current Drupal root.",
	)
	module.DefineStringFlag(
		"root",
		"/path/to/drupal",
		"A full path to a Drupal install directory.",
	)

	cmd.DefineSubCommand("site", "Creates a new site skel.", newSite, "name")

	return cmd
}

func newModule(c cli.Command) {
	var name = c.Param("name")

	rootFlag := fmt.Sprintf("%s", c.Flag("root"))

	if rootFlag == "/path/to/drupal" {

		var err error
		rootFlag, err = os.Getwd()

		if err != nil {
			fmt.Println("Could not get the current directory.")
			return
		}
	} else {
		rootFlag = path.Clean(rootFlag)

		if strings.Contains(rootFlag, "~") {
			fmt.Println("Root must be a fully specified path. E.g. it should not contain a \"~\".")
			return
		}
	}


	root, err := drupalRoot(rootFlag)

	if err != nil {
		fmt.Println(err)
		fmt.Println("You may want to try again with the --root flag.")
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
