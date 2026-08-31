package main

import (
	"archive/tar"
	"bufio"
	"compress/bzip2"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// 1. download the files
// 1.1 if it's compressed uncompress them
// 2. Compute concurrently sha256 signatures of these files
// 3. See if they math the ones in the index file. The index file is sha256sum.txt
// 4.1 Print the number of processed files
// 4.2 If there's a mismatch, print the offending file(s) and exit the program with non-zero value

func main() {
	URL := "https://storage.googleapis.com/353solutions/c/data/taxi.tar"
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	filesPath := filepath.Join(homeDir, "Downloads", "taxi.tar")
	start := time.Now()
	if err := DownloadFiles(URL, filesPath); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("✅ Files downloaded in %v: %s\n", time.Since(start), filesPath)
	}

	unzipDestination := filepath.Join(homeDir, "Downloads", "taxi")
	if err := UnzipTar(filesPath, unzipDestination); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("✅ Unzipped files to:", unzipDestination)
	}

	unzipDestination = filepath.Join(unzipDestination, "taxi")

	sigFile, err := os.Open(filepath.Join(unzipDestination, "sha256sum.txt"))
	if err != nil {
		fmt.Println("Failed to read sha256sum.txt in dir:", err)
	}
	defer sigFile.Close()

	scanner := bufio.NewScanner(sigFile)
	ok := true
	var wg sync.WaitGroup
	matchedFiles := []string{}
	start = time.Now()
	fmt.Println("⌛️ Parsing hashes...")
	for ok {
		hash, file, err := parseSigFile(scanner)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break // Gracefully stop when end of file is reached
			}
			fmt.Println("Error parsing line:", err)
			break
		}

		wg.Add(1)
		go func(path, expected string) {
			defer wg.Done()
			calculatedHash, err := fileSig(path)
			if err != nil {
				fmt.Println(err)
				return
			} else if calculatedHash != expected {
				fmt.Println("❌ Hash mismatch!", path)
				fmt.Println(calculatedHash, "!=", expected)
				ok = false
			} else {
				matchedFiles = append(matchedFiles, path)
			}
		}(filepath.Join(unzipDestination, file+".bz2"), hash)
	}
	wg.Wait()
	if ok {
		fmt.Println("All hashes matched! ✅")
	}
	fmt.Printf("Matched %v files in %v\n", len(matchedFiles), time.Since(start))

}

func DownloadFiles(url, destination string) error { // returns path to download
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Get request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Get failed: %s", resp.Status)
	}

	out, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf("Failed to write: %w", err)
	}
	return nil

}

func UnzipTar(zipPath, destination string) error {
	tarFile, err := os.Open(zipPath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer tarFile.Close()

	tr := tar.NewReader(tarFile)

	// create subfolder to unzip into
	if err := os.MkdirAll(destination, 0755); err != nil {
		return fmt.Errorf("Failed to create sub-folder: %w", err)
	}

	for {
		// fetch next file in archive
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("Failed to fetch archive: %w", err)
		}

		fileName := path.Join(destination, header.Name)

		switch header.Typeflag {
		// if compressed item is dir, create dir
		case tar.TypeDir:
			if err := os.MkdirAll(fileName, 0755); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
		// if compressed item is file, create file (and all the parents), and copy into
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(fileName), 0755); err != nil {
				return fmt.Errorf("failed to create parent directory: %w", err)
			}
			out, err := os.Create(fileName)
			if err != nil {
				return fmt.Errorf("Failed to create file to unpack into: %w", err)
			}
			defer out.Close()

			// Write to file
			if _, err := io.Copy(out, tr); err != nil {
				return fmt.Errorf("Failed to write to file: %w", err)
			}
		}

	}
	return nil

}

func fileSig(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("Failed to open file: %w", err)
	}
	defer file.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, bzip2.NewReader(file))
	if err != nil {
		return "", fmt.Errorf("Failed to write hash: %w", err)
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func parseSigFile(scanner *bufio.Scanner) (string, string, error) {
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		hashAndFile := strings.Fields(line)
		return hashAndFile[0], hashAndFile[1], nil
	}
	return "", "", io.EOF
}
