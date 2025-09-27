package main

import (
    "crypto/sha256"
    "encoding/hex"
    "time"
)

type Block struct {
    Index     int
    Timestamp string
    Data      string
    Hash      string
    PrevHash  string
}

func CreateGenesisBlock() Block {
    return Block{
        Index:     0,
        Timestamp: time.Now().String(),
        Data:      "Genesis Block",
        Hash:      "",
        PrevHash:  "",
    }
}

func GenerateBlock(prevBlock Block, data string) Block {
    newBlock := Block{
        Index:     prevBlock.Index + 1,
        Timestamp: time.Now().String(),
        Data:      data,
        PrevHash:  prevBlock.Hash,
    }
    newBlock.Hash = CalculateHash(newBlock)
    return newBlock
}

func CalculateHash(block Block) string {
    record := string(block.Index) + block.Timestamp + block.Data + block.PrevHash
    h := sha256.New()
    h.Write([]byte(record))
    return hex.EncodeToString(h.Sum(nil))
}

func IsBlockValid(newBlock, prevBlock Block) bool {
    if prevBlock.Index+1 != newBlock.Index {
        return false
    }
    if prevBlock.Hash != newBlock.PrevHash {
        return false
    }
    if CalculateHash(newBlock) != newBlock.Hash {
        return false
    }
    return true
}
