package main

import (
	"fmt"
	"os"
	"os/exec"
	"src/tools"
	"strings"
)

func main() {
	fmt.Println("Started testing " + tools.LibraryFilePath)

	libraryBytes, err := os.ReadFile(tools.LibraryFilePath)
	if err != nil {
		panic(fmt.Errorf("failed to read library file: %w", err))
	}

	var testFileName = strings.Replace(tools.LibraryFilePath, ".template", "_test.go", 1)

	err = createTestFile(testFileName, string(libraryBytes))
	if err != nil {
		panic(fmt.Errorf("failed to create go test file: %w", err))
	}

	err = runGoTests()
	if err != nil {
		panic(fmt.Errorf("tests failed: %w", err))
	}

	err = removeTestFile(err, testFileName)
	if err != nil {
		panic(fmt.Errorf("failed to delete test file: %w", err))
	}
}

const header = `package main

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLeetcodeLibrary(t *testing.T) {
	// We need to check that the template builds.
	// If it has some syntax errors, go test will fail
	assert.True(t, true)
}

`

func createTestFile(testFileName, libraryContent string) error {
	testFile, err := os.Create(testFileName)
	if err != nil {
		return err
	}

	_, err = testFile.WriteString(header + libraryContent)
	if err != nil {
		return nil
	}

	fmt.Println("Created test file")
	return nil
}

func runGoTests() error {
	output, err := exec.Command("go", "test", "-v", "./build/...").Output()
	fmt.Print(string(output))
	if err != nil {
		return err
	}

	fmt.Println("All the tests succeeded")
	return nil
}

func removeTestFile(err error, testFileName string) error {
	err = os.Remove(testFileName)
	if err != nil {
		return err
	}

	fmt.Println("Removed test file")
	return nil
}
