package webAscii

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(file string) error {
	url := ""
	switch file {
	case "public/standard.txt":
		url = "https://github.com/kherldhussein/asciiart/raw/master/standard.txt"
	case "public/shadow.txt":
		url = "https://github.com/kherldhussein/asciiart/raw/master/shadow.txt"
	case "public/thinkertoy.txt":
		url = "https://github.com/kherldhussein/asciiart/raw/master/thinkertoy.txt"
	default:
		return fmt.Errorf("unsupported file name: %s", file)
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Write to file
	err = os.WriteFile(file, body, 0o644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	fmt.Println("Downloaded", file, "from", url)
	return nil
}
