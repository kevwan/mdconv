package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/kevwan/mdconv/converter"
)

var (
	input      = flag.String("i", "", "markdown file")
	output     = flag.String("o", "", "output html file")
	headerFile = flag.String("header", "templates/header.html", "the header template file")
	footerFile = flag.String("footer", "templates/footer.html", "the footer template file")
)

func main() {
	flag.Parse()

	if len(*input) == 0 {
		flag.Usage()
	}

	content, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Fatal(err)
	}

	html, err := converter.MarkdownToHtml(*headerFile, *footerFile, content)
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
