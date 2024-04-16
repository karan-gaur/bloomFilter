package common

import (
	"bloomf/bloom"
	"log"
)

/* Verify if the word found in a file by Dory indeed exists. */
func verifyWordInFile(fileNum uint, keyword string) bool {
	// Open the file
    fileNumber := strconv.FormatUint(uint64(fileNum), 10)
	file, err := os.Open(filepath.Join("sample_docs", fileNumber))
	if err != nil {
        log.Fatal(err)
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

func checkFileExists(filePath string) {
	if _, err := os.Stat(filePath); err == nil {
        log.Printf("File '%s' exists\n", filePath)
		return true
    } else if os.IsNotExist(err) {
        log.Printf("File '%s' does not exist", filePath)
		return false
    } else {
        log.Printf("Error: %v\n", err)
		panic(err)
    }
}

func deleteFile(filePath string) {
	if checkFileExists(filePath) {
        err := os.Remove(filePath)
        if err != nil {
            log.Printf("Error deleting file: %v\n", err)
            return
        }
        log.Printf("Successfully Deleted file - '%s'\n", filePath)
	} else {
		log.Printf("No such file - '%s'", filePath)
	}
}

func createFile(filePath string) {
	// Check if the file exists
	if checkFileExists(filePath) {
        deleteOutputFile()
	}

    file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file: %v\n", err)
		panic(err)
	}
	defer file.Close()
	log.Printf("Created file - '%s'\n", filePath)
}

func createBloomFilter()