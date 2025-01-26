package rpc

import (
	"context"
	"net/http"
	"time"

	"github.com/dwdwow/golimiter"
)

type Client struct {
	URL            string
	Header         http.Header
	Limiter        *golimiter.ReqLimiter
	PostErrHandler PostErrHandler
}

func New(endpoint string) *Client {
	if endpoint == "" {
		endpoint = SOLANA_MAINNET_RPC_URL
	}
	clt := &Client{URL: endpoint}
	clt.Limiter = golimiter.NewReqLimiter(time.Second, 1)
	clt.PostErrHandler = DefaultPostErrHandler
	clt.Header = http.Header{}
	clt.Header.Set("Content-Type", "application/json")
	return clt
}

func (c *Client) GetAccountInfo(address string, opts *GetAccountInfoOptions) (*GetAccountInfoResult, error) {
	return Post[*GetAccountInfoResult](context.Background(), c, "getAccountInfo", address, opts)
}

func (c *Client) GetBalance(address string, opts *GetBalanceOptions) (*GetBalanceResult, error) {
	return Post[*GetBalanceResult](context.Background(), c, "getBalance", address, opts)
}

func (c *Client) GetParsedBlock(slot uint64, opts *GetParsedBlockOptions) (*GetParsedBlockResult, error) {
	_opts := &GetBlockOptions{Encoding: EncodingJsonParsed, MaxSupportedTransactionVersion: maxSupportedTransactionVersion}
	if opts != nil {
		_opts.Commitment = opts.Commitment
		_opts.TransactionDetails = opts.TransactionDetails
		_opts.Rewards = opts.Rewards
	}
	return Post[*GetParsedBlockResult](context.Background(), c, "getBlock", slot, _opts)
}

func (c *Client) GetParsedTransaction(signature string, opts *GetParsedTransactionOptions) (*GetParsedTransactionResult, error) {
	_opts := &GetTransactionOptions{Encoding: EncodingJsonParsed, MaxSupportedTransactionVersion: maxSupportedTransactionVersion}
	if opts != nil {
		_opts.Commitment = opts.Commitment
	}
	return Post[*GetParsedTransactionResult](context.Background(), c, "getTransaction", signature, _opts)
}
