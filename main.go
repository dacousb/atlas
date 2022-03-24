package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options...] [url...]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  [url...]\n")
		fmt.Fprintf(os.Stderr, "        (mandatory) URL(s) of the file(s) to download\n")

		flag.PrintDefaults()
	}

	out := flag.String("o", "", "(optional) name of the output file")
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if len(flag.Args()) > 1 {
		if *out != "" {
			warn("ignoring the -o flag because multiple URLs have been specified")
		}

		for _, url := range flag.Args() {
			fmt.Println(url)
			f := getFile(url)

			if err := os.Rename(f, path.Base(url)); err != nil {
				panic(err)
			}
		}
	} else {
		url := flag.Arg(0)
		f := getFile(url)

		if *out == "" {
			*out = path.Base(url) // get the filename from the url
		}
		if err := os.Rename(f, *out); err != nil {
			panic(err)
		}
	}
}
