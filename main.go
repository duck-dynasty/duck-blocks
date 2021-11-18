package main

import "fmt"

func main() {
	fmt.Println("START")
	b := NewBlockchain()

	t1 := Transaction{
		Sender:    "tom",
		Amount:    100,
		Recipient: "Bob",
	}

	t2 := Transaction{
		Sender:    "Bob",
		Amount:    50,
		Recipient: "Alice",
	}

	t3 := Transaction{
		Sender:    "Alice",
		Amount:    10,
		Recipient: "Tom",
	}

	index := b.NewTransaction(t1)
	fmt.Printf("\nA new transaction will be added to Block %d", index)
	index = b.NewTransaction(t2)
	fmt.Printf("\nA new transaction will be added to Block %d", index)

	fmt.Println("\n")
	fmt.Println("CURRENT CHAIN")
	fmt.Println(b.Chain)

	fmt.Println("")
	fmt.Println("MINE")
	lastBlock := b.LastBlock()
	lastProof := lastBlock.Proof
	proof := b.ProofOfWork(lastProof)
	previousHash := Hash(lastBlock)
	_ = b.NewBlock(proof, previousHash)
	fmt.Println(b.Chain)

	index = b.NewTransaction(t3)
	fmt.Printf("\nA new transaction will be added to Block %d", index)
	fmt.Println("")
	fmt.Println("MINE")
	lastBlock = b.LastBlock()
	lastProof = lastBlock.Proof
	proof = b.ProofOfWork(lastProof)
	previousHash = Hash(lastBlock)
	_ = b.NewBlock(proof, previousHash)
	fmt.Println(b.Chain)

}
