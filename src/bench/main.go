package main

import (
	"log"
	"flag"
    "bufio"
	"bloomf/src/common"
	"bloomf/src/bloom/bloomfilter"
)

func addKeywords(filePath string, bf *bloomfilter.BloomFilter) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
		panic(err)
	}
	defer file.close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		bf.add(scanner.Text())
	}

	// Check for any scanner errors
	if err := scanner.Err(); err != nil {
		log.Printf("Scanner error: %v\n", err)
		panic(err)
	}
}

func createBloomFilter(dataFile string, numDocs int, bf_sz int, hashFunctions int) {
	bloomFilter := make([]bloomfilter.BloomFilter, numDocs)

    files,_ := ioutil.ReadDir(inputDir)
	if len(files) < numDocs {
		numDocs = len(files)
	}

	for i := 0 ; i < numDocs ; i++ {
		bloomFilter[i] = bloomfilter.NewBloomFilter(bf_sz, hashFunctions)


	} 
}

func runInteractiveSearches(dataFile string, numDocs int, bfSize int, hashFunctions int) {
	bf = 
}

func runFileSearch(dataFile string, keyword string, numDocs int, bfSize int, hashFunctions int) {

}

func main() {
	data := flag.String("data", "sample_docs", "Folder with data for Bloom Filter")
	fileSearch := flag.Bool("file_search", false, "Interactive search with terminal or through file")
	keyword := flag.String("keywords", "keywords.txt", "Keywords containts words to search in Bloom Filter")
	numDocs := flag.Int("num_docs", 128, "Max number of Docs to be considered for searching")
	bfSize := flag.Int("bf_sz", 1024, "Size of Bloom Filter")
	hashFunctions := flag.Int("hash_lvl", 5, "Max number of hash functions for Bloom Filter")
	flag.Parse()

	if (fileSearch) {
		runFileSearch(*data, *keyword, *numDocs, *bfSize, *hashFunctions)
	} else {
		runInteractiveSearches(*data, *numDocs, *bfSize, *hashFunctions)
	}
}