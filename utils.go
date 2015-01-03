package main

import (
	"errors"
	"os"
	"path"
)

func drupalRoot(curr string) (root string, err error) {

	var _, fileErr = os.Stat(curr + "/includes/bootstrap.inc")

	if fileErr != nil {
		dir := path.Dir(curr)

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
