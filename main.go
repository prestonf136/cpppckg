package main

import (
	"flag"
	"regexp"

	"./lib"
)

func main() {

	var ghvar bool
	flag.BoolVar(&ghvar, "gh", false, "tells the package manager to use github")
	flag.Parse()

	var args []string = flag.Args()

	if ghvar && len(args) >= 2 {
		match, _ := regexp.MatchString("https://github.com", args[0])

		switch match {
		case true:
			{
				lib.Handleurl(args[0], args[1])
			}
		case false:
			{
				lib.Handlename(args[0], args[1])
			}

		}
	}
}
