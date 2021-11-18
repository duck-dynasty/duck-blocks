package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"time"
)

type Transaction struct {
	Sender    string  `json:"sender"`
	Amount    float32 `json:"amount"`
	Recipient string  `json:"recipient"`
}

type Block struct {
	Index        int           `json:"index"`
	Proof        string        `json:"proof"`
	Timestamp    int64         `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	PreviousHash string        `json:"previous_hash"`
}

type Blockchain struct {
	Chain               []Block       `json:"chain"`
	CurrentTransactions []Transaction `json:"current_transactions"`
}

func (b *Blockchain) LastBlock() Block {
	return b.Chain[len(b.Chain)-1]
}

func (b *Blockchain) NewTransaction(newTransaction Transaction) int {
	b.CurrentTransactions = append(b.CurrentTransactions, newTransaction)

	return b.LastBlock().Index + 1
}

func (b *Blockchain) NewBlock(proof string, previousHash string) Block {
	block := Block{
		Index:        len(b.Chain) + 1,
		Proof:        proof,
		Timestamp:    time.Now().Unix(),
		Transactions: b.CurrentTransactions,
		PreviousHash: previousHash,
	}

	b.CurrentTransactions = nil
	b.Chain = append(b.Chain, block)

	return block
}

func Hash(block Block) string {
	s, _ := json.Marshal(block)
	s2 := md5.Sum(s)

	d := fmt.Sprintf("%x", s2)

	return d
}

func (b *Blockchain) ProofOfWork(lastProof string) string {
	proof := 0

	for !ValidProof(lastProof, string(proof)) {
		proof += 1
	}

	return fmt.Sprintf("%d", proof)
}

func ValidProof(lastProof string, proof string) bool {
	guess := fmt.Sprintf("%s%s", lastProof, proof)
	guessHash := md5.Sum([]byte(guess))
	guessHex := fmt.Sprintf("%x", guessHash)

	if string(guessHex[0]) == "0" {
		return true
	} else {
		return false
	}
}

func NewBlockchain() Blockchain {
	blockchain := Blockchain{
		Chain:               []Block{},
		CurrentTransactions: []Transaction{},
	}

	_ = blockchain.NewBlock("100", "1")

	return blockchain
}
