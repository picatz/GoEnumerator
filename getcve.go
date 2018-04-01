package main

import (
	"compress/gzip"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func getCVE(Year string) {

	fileURL := "https://nvd.nist.gov/feeds/json/cve/1.0/nvdcve-1.0-" + Year + ".json.gz"
	targetDIR := "CVE/"
	Gfile := "CVE/nvdcve-1.0-" + Year + ".json.gz"
	err := DownloadFile(Gfile, fileURL)
	if err != nil {
		panic(err)
	}

	fmt.Println(Gile)
	error := DecompressFile(Gfile, targetDIR)
	if err != nil {
		panic(error)
	}

}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// DecompressFile  here we decompress downloaded gzip files
func DecompressFile(FileName string, target string) error {

	reader, err := os.Open(FileName)
	if err != nil {
		return err
	}
	defer reader.Close()

	archive, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer archive.Close()

	target = filepath.Join(target, archive.Name)
	fmt.Prinln(target)
	writer, err := os.Create(target)
	if err != nil {
		return err
	}
	defer writer.Close()

	_, err = io.Copy(writer, archive)
	return err
}
