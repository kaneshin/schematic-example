package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kaneshin/schematic"
)

var output = flag.String("o", "", "Ouput file")

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()

	log.SetFlags(0)
	log.SetPrefix("schematic: ")

	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatal("missing schema file")
	}

	var i io.Reader
	var err error
	if flag.Arg(0) == "-" {
		i = os.Stdin
	} else {
		if i, err = os.Open(flag.Arg(0)); err != nil {
			log.Fatal(err)
		}
	}

	var o io.Writer
	if *output == "" {
		o = os.Stdout
	} else {
		if o, err = os.Create(*output); err != nil {
			log.Fatal(err)
		}
	}

	var s schematic.Schema
	d := json.NewDecoder(i)
	if err := d.Decode(&s); err != nil {
		log.Fatal(err)
	}

	code, err := Generate(&s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(o, string(code))
}
