package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

type Opt struct {
	out    string
	header bool
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options...] [url...]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  [url...]\n")
		fmt.Fprintf(os.Stderr, "        (mandatory) URL(s) of the file(s) to download\n")

		flag.PrintDefaults()
	}

	var opt Opt
	flag.StringVar(&opt.out, "o", "", "(optional) name of the output file")
	flag.BoolVar(&opt.header, "h", false, "(optional) print the response headers")
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if len(flag.Args()) > 1 {
		if opt.out != "" {
			warn("ignoring the -o flag because multiple URLs have been specified")
		}

		for _, url := range flag.Args() {
			fmt.Println(url)
			f := getFile(url, opt)

			if err := os.Rename(f, path.Base(url)); err != nil {
				panic(err)
			}
		}
	} else {
		url := flag.Arg(0)
		f := getFile(url, opt)

		if opt.out == "" {
			opt.out = path.Base(url) // get the filename from the url
		}
		if err := os.Rename(f, opt.out); err != nil {
			panic(err)
		}
	}
}
