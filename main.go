package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"gitlab.com/golang-commonmark/markdown"
)

var re = regexp.MustCompile(`^\s*#\s*\+kubectl\s*\n`)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please provide a filename")
		os.Exit(1)
	}

	source, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("could not read file: %v", err)
		os.Exit(1)
	}
	md := markdown.New(markdown.XHTMLOutput(true))
	tokens := md.Parse(source)
	code := []string{}
	for _, t := range tokens {
		if t.Tag() == "code" && t.Block() {
			content := t.(*markdown.Fence).Content
			if re.MatchString(content) {
				code = append(code, content, "---")

			}
		}
	}

	fmt.Println(strings.Join(code, "\n"))
}
