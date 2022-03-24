package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func tempFile() *os.File {
	w, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := ioutil.TempFile(w, "atlas")
	if err != nil {
		panic(err)
	}
	return f
}

func getFile(url string) string {
	f := tempFile()
	defer f.Close()

	res, err := http.Get(url)
	if err != nil {
		f.Close()
		exit(err)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		warn("request returned 404")
	}

	p := Progress{total: int(res.ContentLength)}
	if _, err := io.Copy(f, io.TeeReader(res.Body, &p)); err != nil {
		panic(err)
	}
	fmt.Println()

	return f.Name()
}
