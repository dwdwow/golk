package rpc

import (
	"bytes"
	"encoding/json"
	"io"
)

type Req struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	ID      int64  `json:"id"`
}

func NewReqData(method string, params ...any) *Req {
	return &Req{
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
		ID:      1,
	}
}

func (r *Req) ToReader() (io.Reader, error) {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(r); err != nil {
		return nil, err
	}
	return buf, nil
}

type Resp[D any] struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int64  `json:"id"`
	Result  D      `json:"result"`
}

func (r *Resp[D]) FromReader(reader io.Reader) error {
	return json.NewDecoder(reader).Decode(r)
}

// =============================== transaction ===============================

type Transaction struct {
	Signatures []string `json:"signatures"`
}

type TxMessage struct {
	Header          TxMsgHeader `json:"header"`
	AccountKeys     []string    `json:"accountKeys"`
	RecentBlockhash string      `json:"recentBlockhash"`
}

type TxMsgHeader struct {
	NumRequiredSignatures       int64 `json:"numRequiredSignatures"`
	NumReadonlySignedAccounts   int64 `json:"numReadonlySignedAccounts"`
	NumReadonlyUnsignedAccounts int64 `json:"numReadonlyUnsignedAccounts"`
}

type Instruction struct {
	ProgramID string `json:"programId"`
	Data      string `json:"data"`
}

// =============================== getBalance ===============================

type GetBalanceOptions struct {
	Commitment     Commitment `json:"commitment,omitempty"`
	MinContextSlot int64      `json:"minContextSlot,omitempty"`
}

type GetBalanceResult struct {
	Value   uint64 `json:"value"`
	Context struct {
		Slot int64 `json:"slot"`
	} `json:"context"`
}

// =============================== getBlock ===============================

type GetBlockOptions struct {
	Commitment                     Commitment         `json:"commitment,omitempty"`
	Encoding                       Encoding           `json:"encoding,omitempty"`
	MaxSupportedTransactionVersion TransactionVersion `json:"maxSupportedTransactionVersion"`
	TransactionDetails             TransactionDetails `json:"transactionDetails,omitempty"`
	Rewards                        bool               `json:"rewards,omitempty"`
}

type GetParsedBlockOptions struct {
	Commitment                     Commitment         `json:"commitment,omitempty"`
	MaxSupportedTransactionVersion TransactionVersion `json:"maxSupportedTransactionVersion"`
	TransactionDetails             TransactionDetails `json:"transactionDetails,omitempty"`
	Rewards                        bool               `json:"rewards,omitempty"`
}

type GetParsedBlockResult struct {
	// The blockhash of this block.
	Blockhash string `json:"blockhash"`

	// The blockhash of this block's parent;
	// if the parent block is not available due to ledger cleanup,
	// this field will return "11111111111111111111111111111111".
	PreviousBlockhash string `json:"previousBlockhash"`

	// The slot index of this block's parent.
	ParentSlot uint64 `json:"parentSlot"`

	// Present if "full" transaction details are requested.
	Transactions []ParsedTransactionWithMeta `json:"transactions"`

	// Present if "signatures" are requested for transaction details;
	// an array of signatures, corresponding to the transaction order in the block.
	Signatures []string `json:"signatures"`

	// Present if rewards are requested.
	Rewards []BlockReward `json:"rewards"`

	// Estimated production time, as Unix timestamp (seconds since the Unix epoch).
	// Nil if not available.
	BlockTime int64 `json:"blockTime"`

	// The number of blocks beneath this block.
	BlockHeight *uint64 `json:"blockHeight"`
}

// =============================== getTransaction ===============================

type GetTransactionOptions struct {
	Commitment                     Commitment         `json:"commitment,omitempty"`
	Encoding                       Encoding           `json:"encoding,omitempty"`
	MaxSupportedTransactionVersion TransactionVersion `json:"maxSupportedTransactionVersion"`
}

type GetParsedTransactionOptions struct {
	Commitment                     Commitment         `json:"commitment,omitempty"`
	MaxSupportedTransactionVersion TransactionVersion `json:"maxSupportedTransactionVersion"`
}

type GetParsedTransactionResult struct {
	Slot        uint64                 `json:"slot"`
	BlockTime   int64                  `json:"blockTime"`
	Transaction *ParsedTransaction     `json:"transaction"`
	Meta        *ParsedTransactionMeta `json:"meta"`
	Version     int64                  `json:"version"`
}
