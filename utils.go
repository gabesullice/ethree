package main

import (
	"errors"
	"os"
	p "path"
)

func drupalRoot(path ...string) (root string, err error) {
	var curr string

	if len(path) == 0 {
		var getwdErr error

		curr, getwdErr = os.Getwd()

		if getwdErr != nil {
			return "", errors.New("Could not get the current directory.\n")
		}
	} else {
		curr = path[0]
	}

	var _, fileErr = os.Stat(curr + "/includes/bootstrap.inc")

	if fileErr != nil {
		dir := p.Dir(curr)

		if dir == "/" {
			err = errors.New("Not inside of a Drupal installation.\n")
			return "", err
		}

		root, err = drupalRoot(dir)
	} else {
		root = curr
	}

	return root, err
}
