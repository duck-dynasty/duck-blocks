package main

type Node struct{}

type Transaction struct {
	Sender    string  `json:"sender"`
	Amount    float32 `json:"amount"`
	Recipient string  `json:"recipient"`
}

type Message struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
	Chat   string `json:"chat"`
}

type Block struct {
	Index        int           `json:"index"`
	Proof        string        `json:"proof"`
	Timestamp    int64         `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	PreviousHash string        `json:"previous_hash"`
}

type Blockchain struct {
	Chain               []Block         `json:"chain"`
	CurrentTransactions []Transaction   `json:"current_transactions"`
	Nodes               map[string]Node `json:"nodes"`
}

func (b *Blockchain) LastBlock() Block {
	return b.Chain[len(b.Chain)-1]
}
