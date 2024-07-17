package webAscii

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	send "webAscii/utils"
)

var expectedChecksum = map[string]string{
	"public/standard.txt":   "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
	"public/shadow.txt":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
	"public/thinkertoy.txt": "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3",
}

func ValidateFileChecksum(w http.ResponseWriter, file string) error {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		send.SendError(w, fmt.Sprintf("Error 500: Please wait while downloading... %v", err), http.StatusInternalServerError)
		err := DownloadFile(file)
		if err != nil {
			return fmt.Errorf("error downloading file: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("error checking file existence: %w", err)
	} else {
		// File exists, calculate its checksum
		checksum, err := calculateChecksum(file)
		if err != nil {
			send.SendError(w, fmt.Sprintf("Error 500: Internal server error: %v", err), http.StatusInternalServerError)
			return fmt.Errorf("no expected checksum defined for file: %s", file)
		}

		expected, ok := expectedChecksum[file]
		if !ok {
			send.SendError(w, fmt.Sprintf("Error 500: Internal server error: %v", err), http.StatusInternalServerError)
			return fmt.Errorf("no expected checksum defined for file: %s", file)
		}

		if checksum != expected {
			send.SendError(w, fmt.Sprintf("Error 500: Internal server error: %v", err), http.StatusInternalServerError)
			err := DownloadFile(file)
			if err != nil {
				return fmt.Errorf("error downloading file: %w", err)
			}
		}

		// fmt.Printf("Checksum verified for file %s\n", file)
	}
	return nil
}

func calculateChecksum(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
