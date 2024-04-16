package bloom

import (
    "crypto/sha256"
)

type BloomFilter struct {
    size           int
    hashFunctions  int
    bitArray       []bool
}

func NewBloomFilter(size, hashFunctions int) *BloomFilter {
    return &BloomFilter{
        size:          size,
        hashFunctions: hashFunctions,
        bitArray:      make([]bool, size),
    }
}

func (bf *BloomFilter) Add(item string) {
    for i := 0; i < bf.hashFunctions; i++ {
        index := int(bf.getHash(item, i)) % bf.size
        bf.bitArray[index] = true
    }
}

func (bf *BloomFilter) Contains(item string) bool {
    for i := 0; i < bf.hashFunctions; i++ {
        index := int(bf.getHash(item, i)) % bf.size
        if index >= bf.size || !bf.bitArray[index] {
            return false
        }
    }
    return true
}

func (bf *BloomFilter) getHash(item string, i int) uint32 {
    hash := sha256.New()
    hash.Write([]byte(item + string(i)))
    hashSum := hash.Sum(nil)
    return uint32(hashSum[0]) | uint32(hashSum[1])<<8 | uint32(hashSum[2])<<16 | uint32(hashSum[3])<<24
}