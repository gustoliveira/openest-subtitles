package cmd

import (
	"encoding/xml"
	"fmt"
	"os"
)

type NfoFile struct {
	ImdbID    string `xml:"imdbid"`
	Title     string `xml:"title"`
	Year      string `xml:"year"`
	ShowTitle string `xml:"showtitle"`
	Episode   string `xml:"episode"`
	Season    string `xml:"season"`
}

func GetNfoFile(path string) (NfoFile, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return NfoFile{}, err
	}

	return parseNfoFile(fileBytes)
}

func parseNfoFile(fileBytes []byte) (NfoFile, error) {
	var nfoFile NfoFile

	err := xml.Unmarshal(fileBytes, &nfoFile)
	if err != nil {
		return NfoFile{}, err
	}

	return nfoFile, nil
}
