package util

import (
	"archive/zip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/jszwec/csvutil"
)

func WriteCsvFile(filename string, v interface{}) error {
	b, err := csvutil.Marshal(v)
	if err != nil {
		return err
	}
	ioutil.WriteFile(filename, b, 0666)
	return nil
}

func ZipBase64Encode(filename string, files []string) (string, error) {

	err := ZipFiles(filename, files)
	if err != nil {
		return "", err
	}
	data, _ := ioutil.ReadFile(filename)

	return base64.StdEncoding.EncodeToString(data), nil
}

// ZipFiles compresses one or many files into a single zip archive file.
// Param 1: filename is the output zip file's name.
// Param 2: files is a list of files to add to the zip.
func ZipFiles(filename string, files []string) error {

	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}

	zipWriter := zip.NewWriter(newZipFile)
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Trapped panic: %s (%T)\n", err, err)
		}
		newZipFile.Close()
		zipWriter.Close()
	}()

	// Add files to zip
	for _, file := range files {
		if err = addFileToZip(zipWriter, file); err != nil {
			return err
		}
	}
	return nil
}

func addFileToZip(zipWriter *zip.Writer, filename string) error {

	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// Using FileInfoHeader() above only uses the basename of the file. If we want
	// to preserve the folder structure we can overwrite this with the full path.
	header.Name = filename

	// Change to deflate to gain better compression
	// see http://golang.org/pkg/archive/zip/#pkg-constants
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
