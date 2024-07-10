package webAscii

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	StandardAsciiURL   = "https://github.com/kherldhussein/asciiart/raw/master/standard.txt"
	ShadowAsciiURL     = "https://github.com/kherldhussein/asciiart/raw/master/shadow.txt"
	ThinkertoyAsciiURL = "https://github.com/kherldhussein/asciiart/raw/master/thinkertoy.txt"
)

var fileURLs = map[string]string{
	"public/standard.txt":   StandardAsciiURL,
	"public/shadow.txt":     ShadowAsciiURL,
	"public/thinkertoy.txt": ThinkertoyAsciiURL,
}

func DownloadFile(file string) error {
	url, ok := fileURLs[file]
	if !ok {
		return fmt.Errorf("unsupported file name: %s", file)
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch URL %s: %w", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	err = os.WriteFile(file, body, 0o644)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %w", file, err)
	}

	fmt.Printf("Downloaded %s from %s\n", file, url)
	return nil
}
