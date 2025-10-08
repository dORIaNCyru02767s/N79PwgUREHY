// 代码生成时间: 2025-10-09 01:58:25
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
)

// BlockchainNode represents a basic structure for a blockchain node
type BlockchainNode struct {
    ID       string
    Peers    []*BlockchainNode
    Transactions []Transaction
    Chain    []Block
}

// Transaction represents a basic structure for a transaction
type Transaction struct {
    ID     string
    From   string
    To     string
    Amount float64
}

// Block represents a basic structure for a block
type Block struct {
    Index     int
    Timestamp string
    Transactions []Transaction
    PrevHash  string
    Hash      string
}

// NewBlockchainNode creates a new blockchain node
func NewBlockchainNode(id string) *BlockchainNode {
    node := &BlockchainNode{ID: id}
    node.Peers = make([]*BlockchainNode, 0)
    node.Transactions = make([]Transaction, 0)
    node.Chain = make([]Block, 0)
    return node
}

// AddTransaction adds a new transaction to the blockchain node's transaction pool
func (n *BlockchainNode) AddTransaction(t Transaction) {
    n.Transactions = append(n.Transactions, t)
}

// AddPeer adds a new peer to the blockchain node
func (n *BlockchainNode) AddPeer(node *BlockchainNode) {
    n.Peers = append(n.Peers, node)
}

// CreateGenesisBlock creates the genesis block
func (n *BlockchainNode) CreateGenesisBlock() {
    genesisBlock := Block{
        Index: 0,
        Timestamp: time.Now().String(),
        Transactions: []Transaction{},
        PrevHash: "0",
        Hash: CalculateHash(genesisBlock),
    }
    n.Chain = append(n.Chain, genesisBlock)
}

// CalculateHash calculates the hash of a block
func CalculateHash(block Block) string {
    // Implement hash calculation logic here
    // For simplicity, this is just a placeholder
    return fmt.Sprintf("%x", block.Index)
}

// StartAPI starts the Fiber API server
func StartAPI(node *BlockchainNode) {
    app := fiber.New()

    // Health check endpoint
    app.Get("/health", func(c *fiber.Ctx) error {
        return c.SendString("Blockchain node is up and running")
    })

    // Add transaction endpoint
    app.Post("/transactions/new", func(c *fiber.Ctx) error {
        var t Transaction
        if err := c.BodyParser(&t); err != nil {
            return err
        }
        node.AddTransaction(t)
        return c.SendString("Transaction added")
    })

    // Get chain endpoint
    app.Get("/chain", func(c *fiber.Ctx) error {
        return c.JSON(node.Chain)
    })

    // Start the server
    if err := app.Listen(":3000"); err != nil {
        fmt.Println("Error starting server: ", err)
    }
}

func main() {
    node := NewBlockchainNode("node1")
    node.CreateGenesisBlock()
    StartAPI(node)
}
