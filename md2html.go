package main

import (
	"embed"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/kevwan/mdconv/converter"
)

const (
	embedDir        = "templates"
	embedHeaderFile = "header.html"
	embedFooterFile = "footer.html"
)

var (
	input      = flag.String("i", "", "markdown file")
	output     = flag.String("o", "", "output html file")
	headerFile = flag.String("header", "", "the header template file, default to use embedded template")
	footerFile = flag.String("footer", "", "the footer template file, default to use embedded template")

	//go:embed templates/*.html
	embedded embed.FS
)

func getContent(file, embeddedFile string) (string, error) {
	if len(file) > 0 {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return "", err
		}

		return string(content), nil
	}

	content, err := embedded.ReadFile(path.Join(embedDir, embeddedFile))
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func main() {
	flag.Parse()

	if len(*input) == 0 {
		flag.Usage()
	}

	header, err := getContent(*headerFile, embedHeaderFile)
	if err != nil {
		log.Fatal(err)
	}

	footer, err := getContent(*footerFile, embedFooterFile)
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Fatal(err)
	}

	html, err := converter.MarkdownToHtml(header, footer, content)
	if err != nil {
		log.Fatal(err)
	}

	if len(*output) > 0 {
		file, err := os.Create(*output)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(file, string(html))
	} else {
		fmt.Println(string(html))
	}
}
