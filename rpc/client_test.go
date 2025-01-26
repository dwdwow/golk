package rpc

import (
	"os"
	"testing"
)

func TestGetBalance(t *testing.T) {
	client := New(os.Getenv("SOL_RPC_HTTP_URL"))
	balance, err := client.GetBalance("9WzDXwBbmkg8ZTbNMqUxvQRAyrZzDsGYdLVL9zYtAWWM", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(balance)
}

func TestGetParsedBlock(t *testing.T) {
	client := New(os.Getenv("SOL_RPC_HTTP_URL"))
	block, err := client.GetParsedBlock(316503550, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("block tx count:", len(block.Transactions))
}

func TestGetParsedTransaction(t *testing.T) {
	client := New(os.Getenv("SOL_RPC_HTTP_URL"))
	tx, err := client.GetParsedTransaction("mCN8EsjzE4HELqF8CBC845wvL39mGDSGoVq6ex8yAGqjxGN59U5MNt7H7rkA4NLgQiJ6bjXefEy45oqN41KmUm4", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tx)
}
