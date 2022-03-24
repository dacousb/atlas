package main

import (
	"fmt"
	"strings"
)

const (
	kb = 1 << 10
	mb = 1 << 20
	gb = 1 << 30
)

type Progress struct {
	curr  int
	total int
}

func getSize(n int) string {
	s := float64(n)
	if s < 0 {
		return "?"
	} else if s/gb > 1 {
		return fmt.Sprintf("%.2f GB", s/gb)
	} else if s/mb > 1 {
		return fmt.Sprintf("%.2f MB", s/mb)
	} else if s/kb > 1 {
		return fmt.Sprintf("%.2f KB", s/kb)
	} else {
		return fmt.Sprintf("%d B", n)
	}
}

func (p *Progress) Write(b []byte) (int, error) {
	c := len(b) // chunk size
	p.curr += c

	fmt.Printf("\r%s\r%s / %s", strings.Repeat(" ", 30),
		getSize(p.curr), getSize(p.total))

	return c, nil
}
