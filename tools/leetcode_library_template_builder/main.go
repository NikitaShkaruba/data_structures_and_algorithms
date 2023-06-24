package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"src/tools"
)

// This file gets the content of all the src files, and concatenates them into one, putting them into a single package.
// The concatenated file should be pasted into leetcode.com and used easily

func main() {
	fmt.Println("Generation started")

	outputFile, err := createOutputFile()
	if err != nil {
		panic(fmt.Errorf("failed to create the output file: %w", err))
	}

	sourceFilePaths, err := getSourceFilePaths()
	if err != nil {
		panic(fmt.Errorf("failed to retrieve source file paths: %w", err))
	}

	err = concatenateSourceFiles(outputFile, sourceFilePaths)
	if err != nil {
		panic(fmt.Errorf("failed to concatenate source files: %w", err))
	}

	fmt.Println("Successfully generated " + tools.LibraryTemplateFilePath)
}

func createOutputFile() (*os.File, error) {
	outputFile, err := os.Create(tools.LibraryTemplateFilePath)
	if err != nil {
		return nil, err
	}

	fmt.Println("Created output file")
	return outputFile, err
}

const header = `////////////////////////// Leetcode library //////////////////////////

//////////////////////////////////////////////////////////////////////
// This is my auto-generated template that I use when leetcoding.   //
//                                                                  //
// Standard go library is pretty weak and inconvenient, so I've     //
// decided to write a good library myself. I copy paste this        //
// template every time I'm when writing a solution and enjoy fully  //
// working data structures with good interfaces in go.              //
//                                                                  //
// Source code:                                                     //
// https://github.com/NikitaShkaruba/data_structures_and_algorithms //
//                                                                  //
// Library usage example:                                           //
// https://leetcode.com/submissions/detail/978392175                //
//////////////////////////////////////////////////////////////////////
`

func concatenateSourceFiles(outputFile *os.File, sourceFilePaths []string) error {
	fmt.Println("Writing output file parts:")

	_, err := outputFile.WriteString(header)
	if err != nil {
		return err
	}
	logOutputFilePart("[Header]")

	for _, path := range sourceFilePaths {
		logOutputFilePart(path)

		libraryPieceBytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		libraryPiece := cutUnnecessaryFromLibraryPiece(string(libraryPieceBytes))

		_, err = outputFile.WriteString(libraryPiece)
		if err != nil {
			return err
		}
	}

	return nil
}

func getSourceFilePaths() ([]string, error) {
	filePaths := make([]string, 0)

	err := filepath.Walk("src", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only pass .go files
		if len(path) < 3 || path[len(path)-3:] != ".go" {
			return nil
		}

		// Only pass not test.go files
		if len(path) > 8 && path[len(path)-8:] == "_test.go" {
			return nil
		}

		// Ignore algorithms folder
		if len(path) > 13 && path[:14] == "src/algorithms" {
			return nil
		}

		filePaths = append(filePaths, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return filePaths, nil
}

func cutUnnecessaryFromLibraryPiece(libraryPiece string) string {
	// Cut down package declarations
	libraryPiece = regexp.MustCompile("package .*?\n").ReplaceAllString(libraryPiece, "")

	// Cut down multiline imports declarations
	libraryPiece = regexp.MustCompile("import \\(\n(\t\".*?\"\n)+\\)\n").ReplaceAllString(libraryPiece, "")

	// Cut down single line import declarations
	libraryPiece = regexp.MustCompile("import \".*?\"\n").ReplaceAllString(libraryPiece, "")

	// Cut down stdlib package
	libraryPiece = regexp.MustCompile("stdlib\\.").ReplaceAllString(libraryPiece, "")

	return libraryPiece
}

func logOutputFilePart(part string) {
	fmt.Printf("- %s\n", part)
}
