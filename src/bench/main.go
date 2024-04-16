package main

import (
	"os"
	"log"
	"flag"
    "bufio"
	"strconv"
    "io/ioutil"
	"path/filepath"
	"bloomf/src/bloom"
	"bloomf/src/common"
)

func addKeywords(filePath string, bf *bloom.BloomFilter) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		bf.Add(scanner.Text())
	}

	// Check for any scanner errors
	if err := scanner.Err(); err != nil {
		log.Printf("Scanner error: %v\n", err)
		panic(err)
	}
}

func createBloomFilter(dataFile string, numDocs int, bf_sz int, hashFunctions int) []*bloom.BloomFilter {

    files,_ := ioutil.ReadDir(dataFile)
	if len(files) < numDocs {
		numDocs = len(files)
	}

	bloomFilter := make([]*bloom.BloomFilter, numDocs)
	for docId, file := range files {
		if docId >= numDocs {
			break
		}
		bloomFilter[docId] = bloom.NewBloomFilter(bf_sz, hashFunctions)
		filePath := filepath.Join(dataFile, file.Name())
		addKeywords(filePath, bloomFilter[docId])
	}
	return bloomFilter
}

func runInteractiveSearches(dataFile string, numDocs int, bfSize int, hashFunctions int) {
	bloomFilters := createBloomFilter(dataFile, numDocs, bfSize, hashFunctions)

	input := bufio.NewScanner(os.Stdin)

    log.Printf("Enter a keyword to search for: ")
    for input.Scan() {
		for i, bf := range bloomFilters {
			if bf.Contains(input.Text()) {
				log.Printf("Word found in file_id - '%d'\n", i)
			} 
		}
        log.Printf("Enter a keyword to search for: ")
	}
}

func runFileSearch(dataFile string, keywords string, numDocs int, bfSize int, hashFunctions int) {
	common.CreateFile("output.csv")
	bloomFilters := createBloomFilter(dataFile, numDocs, bfSize, hashFunctions)

	file, err := os.Open(keywords)
    if err != nil {
        log.Printf("Error reading file: %v", err)
        return
    }
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Storing all False Positives indexex 
		var falsePositives []string
		for i, bf := range bloomFilters {
			if bf.Contains(scanner.Text()) {
				if !common.VerifyWordInFile(dataFile, i, scanner.Text()) {
					// False Positive found for keyword
                    falsePositives = append(falsePositives, strconv.Itoa(i))
				}
			}
		}
		// Adding words to output.csv
		common.AddWordToCsv(scanner.Text(), falsePositives)
	}

}

func main() {
	dataFile := flag.String("data", "sample_docs", "Folder with data for Bloom Filter")
	fileSearch := flag.Bool("file_search", false, "Interactive search with terminal or through file")
	keywords := flag.String("keywords", "keywords.txt", "Keywords containts words to search in Bloom Filter")
	numDocs := flag.Int("num_docs", 128, "Max number of Docs to be considered for searching")
	bfSize := flag.Int("bf_sz", 1024, "Size of Bloom Filter")
	hashFunctions := flag.Int("hash_lvl", 5, "Max number of hash functions for Bloom Filter")
	flag.Parse()

	if *fileSearch {
		runFileSearch(*dataFile, *keywords, *numDocs, *bfSize, *hashFunctions)
	} else {
		runInteractiveSearches(*dataFile, *numDocs, *bfSize, *hashFunctions)
	}
}