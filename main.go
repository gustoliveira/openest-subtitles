package main

import (
	"fmt"
	"os"

	cmd "openest-subtitles/cmd"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide only the .nfo file path. Nothing more or less.")
		return
	}

	nfoFilePath := os.Args[1]

	nfoFile, err := cmd.GetNfoFile(nfoFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("NFO file found: %s\n", nfoFile.ImdbID)
	fmt.Printf("Download Link: %s\n", cmd.GetSubtitleLink(nfoFile.ImdbID))
}
