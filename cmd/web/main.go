package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var (
	errBannerNotFound = errors.New("banner not found")
)

func main() {
	addr := flag.String("addr", ":8000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", home)
	mux.HandleFunc("POST /asciiart", asciiArt)

	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		os.Exit(1)
	}

}

func AsciiartFont(text, banner string) (string, error) {
	file := "banner/" + banner + ".txt"

	data, err := os.ReadFile(strings.TrimSpace(file))
	if err != nil {
		return "", errBannerNotFound
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	font := strings.Split(content, "\n")

	text = strings.ReplaceAll(text, "\r\n", "\n")
	lines := strings.Split(text, "\n")

	var output strings.Builder

	for _, line := range lines {
		if line == "" {
			output.WriteString("\n")
			continue
		}

		for _, char := range line {
			if char < 32 || char > 126 {
				continue
			}
		}
		for row := 0; row < 8; row++ {
			for _, char := range line {
				math := int((char-32)*9) + row + 1
				if math >= 0 && math < len(font) {
					output.WriteString(font[math])
				}
			}
			output.WriteString("\n")
		}
	}
	return output.String(), nil
}

func errorMessage(code int) string {
	return fmt.Sprintf("Error %d", code)
}
