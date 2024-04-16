package common

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"path/filepath"
	"encoding/csv"
)

/* Verify if the word found in a file by Dory indeed exists. */
func verifyWordInFile(fileNum uint, keyword string) bool {
    fileNumber := strconv.FormatUint(uint64(fileNum), 10)
	file, err := os.Open(filepath.Join("sample_docs", fileNumber))
	if err != nil {
        log.Fatal(err)
		panic(err)
    }
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Loop through each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		// Check if the word exists in the line
		if strings.Contains(line, keyword) {
            log.Printf("Found Word in File - %s\n", fileNumber)
            return true
		}
	}

	// Check for scanner errors
	if err != nil {
        log.Fatal(err)
    }
	
    log.Printf("False Positive Word - '%s' in File - %s\n", keyword, fileNumber)
    return false
}

/* Check if given file with FilePath exists */
func checkFileExists(filePath string) bool {
    if _, err := os.Stat(filePath); err == nil {
        return true
    } else if os.IsNotExist(err) {
        return false
    } else {
        log.Printf("Error: %v\n", err)
        panic(err)
    }
}

/* Delte file with given FilePath */
func deleteFile(filePath string) {
	if checkFileExists(filePath) {
        err := os.Remove(filePath)
        if err != nil {
            log.Printf("Error deleting file: %v\n", err)
            return
        }
        log.Printf("Successfully Deleted file - '%s'\n", filePath)
	} else {
		log.Printf("No such file - '%s'\n", filePath)
	}
}

/* Create File by given fileName at given filePath */
func createFile(filePath string) {
	// Check if the file exists
	if checkFileExists(filePath) {
        deleteFile(filePath)
	}

    file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file: %v\n", err)
		panic(err)
	}
	defer file.Close()
	log.Printf("Created file - '%s'\n", filePath)
}

/* Add False Positvie to "output.csv" */
func addWordToCsv(keyword string, numDocs []string) {
    if len(numDocs) == 0 {
        log.Printf("No falsePositives for '%s'\n", keyword)
        return
    }

    filePath := "output.csv"
    if !checkFileExists(filePath) {
		createFile(filePath)
	}

    file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
		panic(err)
	}
	defer file.Close()

    var fields [2]string
    fields[0] = keyword
    fields[1] = strings.Join(numDocs, ";")

    writer := csv.NewWriter(file)
	defer writer.Flush()
    rowSlice := fields[:]
    
	// Write the word to the file
	if err := writer.Write(rowSlice); err != nil {
        log.Fatalf("Error writing to CSV: %v", err)
    }
}
