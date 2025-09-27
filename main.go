package main

import (
    "encoding/json"
    "net/http"
)

var Blockchain []Block

func main() {
    genesisBlock := CreateGenesisBlock()
    Blockchain = append(Blockchain, genesisBlock)

    http.HandleFunc("/blocks", GetBlocksHandler)
    http.HandleFunc("/mine", MineBlockHandler)

    println("Server running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

func GetBlocksHandler(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(Blockchain)
}

func MineBlockHandler(w http.ResponseWriter, r *http.Request) {
    data := r.URL.Query().Get("data")
    newBlock := GenerateBlock(Blockchain[len(Blockchain)-1], data)

    if IsBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
        Blockchain = append(Blockchain, newBlock)
        json.NewEncoder(w).Encode(newBlock)
    } else {
        http.Error(w, "Invalid block", http.StatusBadRequest)
    }
}
