package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convert(md string) string {
	lines := strings.Split(md, "\n")
	line_index := 0
	html := ""

	for true { //while true
		if strings.Contains(lines[line_index], "#") {
			header_count := strings.Count(lines[line_index], "#")
			h1, h2 := fmt.Sprintf("<h%d>", header_count), fmt.Sprintf("</h%d>", header_count)
			html += h1 + strings.TrimLeft(lines[line_index], "#") + h2 + "\n"
		} else if strings.Contains(lines[line_index], "-") {
			html += "<ul>\n"
			for true {
				if strings.Contains(lines[line_index], "-") {
					html += "<li>" + strings.TrimLeft(lines[line_index], "-") + "</li>\n"
				} else {
					break
				}
				line_index++
			}
			html += "</ul>\n"
		}

	}
	return html
}

func main() {
	document_start :=
		`<!DOCTYPE html>
		<html lang="en">
		<head>
		</head>
		<body>`

	document_end :=
		`</body>
		</html>`

	filepath := os.Args[1]

	contents, err := os.ReadFile(filepath)
	check(err)
	html := document_start + convert(string(contents)) + document_end

	f, err := os.Create("output/index.html")
	check(err)
	f.WriteString(html)

	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
