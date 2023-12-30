package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Const struct {
	Name  string
	Value string
}

func main() {
	// Parse command-line arguments
	filePath := flag.String("input", "", "Path to the input file")
	outputPath := flag.String("output", "", "Path to the output file")
	packageName := flag.String("package", "", "Package name for the output file")
	flag.Parse()

	// Open the input file
	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Create the output directory if it does not exist
	outputDir := filepath.Dir(*outputPath)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err = os.MkdirAll(outputDir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Create the output file
	outputFile, err := os.Create(*outputPath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	// Write the package declaration to the output file
	fmt.Fprintf(outputFile, "package %s\n\n", *packageName)

	// Read the input file line by line
	scanner := bufio.NewScanner(file)
	var consts []Const
	maxLen := 0
	for scanner.Scan() {
		line := scanner.Text()

		// Generate a constant declaration for each line
		constName := toConstName(line)
		if len(constName) > maxLen {
			maxLen = len(constName)
		}
		consts = append(consts, Const{Name: constName, Value: line})
	}

	// Write the opening line of the const block
	fmt.Fprintln(outputFile, "const (")

	// Write the constants to the output file
	for _, c := range consts {
		fmt.Fprintf(outputFile, "\t%-*s = \"%s\"\n", maxLen, c.Name, c.Value)
	}

	// Write the closing line of the const block
	fmt.Fprintln(outputFile, ")")

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func toConstName(input string) string {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(input, " ")
	titleSpace := cases.Title(language.English).String(processedString)
	constName := strings.ReplaceAll(titleSpace, " ", "")
	if !unicode.IsLetter(rune(constName[0])) {
		constName = "_" + constName
	}
	return constName
}
