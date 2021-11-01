package converter

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/kennygrant/sanitize"
	"github.com/kevwan/blackfriday"
)

const headerMark = "#"

func Convert(body []byte) []byte {
	toc := buildToc(body)
	content := buildContent(body)
	return append(toc, content...)
}

func MarkdownToHtml(headerFile, footerFile string, body []byte) ([]byte, error) {
	html := Convert(body)

	header, err := ioutil.ReadFile(headerFile)
	if err != nil {
		return nil, err
	}
	footer, err := ioutil.ReadFile(footerFile)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf(string(header), extractTitle(body)))
	buffer.Write(html)
	buffer.Write(footer)

	return buffer.Bytes(), nil
}

func buildToc(body []byte) []byte {
	title := extractTitle(body)
	renderer := blackfriday.HtmlRenderer(blackfriday.HTML_TOC|blackfriday.HTML_OMIT_CONTENTS, title, "")
	return blackfriday.Markdown(body, renderer, getExtensions())
}

func buildContent(body []byte) []byte {
	title := extractTitle(body)
	renderer := blackfriday.HtmlRenderer(blackfriday.HTML_TOC|blackfriday.HTML_OMIT_TOC, title, "")
	return blackfriday.Markdown(body, renderer, getExtensions())
}

func extractTitle(body []byte) string {
	scanner := bufio.NewScanner(bytes.NewReader(body))

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, headerMark) {
			continue
		}

		for {
			if strings.HasPrefix(line, headerMark) {
				line = strings.TrimPrefix(line, headerMark)
			} else {
				break
			}
		}

		return strings.TrimSpace(sanitize.HTML(line))
	}

	return ""
}

func getExtensions() int {
	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS
	return extensions
}

func stripchars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}
