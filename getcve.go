package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func getCVE(Year string) {

	defer wg.Done()
	fileURL := "https://nvd.nist.gov/feeds/json/cve/1.0/nvdcve-1.0-" + Year + ".json.gz"
	targetDIR := "CVE/"
	Gfile := "nvdcve-1.0-" + Year + ".json.gz"
	Uname := "nvdcve-1.0-" + Year + ".json"
	Destination := targetDIR + Gfile
	err := DownloadFile(Destination, fileURL)
	if err != nil {
		panic(err)
	}

	fmt.Println(Gfile)
	error := DecompressFile(Uname, targetDIR, Destination)
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
func DecompressFile(FileName string, target string, Destination string) error {

	reader, err := os.Open(Destination)
	if err != nil {
		return err
	}
	defer reader.Close()

	archive, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer archive.Close()
	fmt.Println(target)
	target = filepath.Join(target, FileName)
	fmt.Println(target)
	fmt.Println(archive.Name)
	writer, err := os.Create(target)
	if err != nil {
		return err
	}
	defer writer.Close()

	_, err = io.Copy(writer, archive)
	return err
}
