package rpc

type ParsedMessageAccount struct {
	PublicKey string `json:"pubkey"`
	Signer    bool   `json:"signer"`
	Writable  bool   `json:"writable"`
}

type ParsedInstructionInfo struct {
	Info            map[string]any `json:"info"`
	InstructionType string         `json:"type"`
}

type ParsedInstruction struct {
	Program     string                `json:"program,omitempty"`
	ProgramId   string                `json:"programId,omitempty"`
	Parsed      ParsedInstructionInfo `json:"parsed,omitempty"`
	Data        string                `json:"data,omitempty"`
	Accounts    []string              `json:"accounts,omitempty"`
	StackHeight int64                 `json:"stackHeight"`
}

type ParsedInnerInstruction struct {
	Index        int64               `json:"index"`
	Instructions []ParsedInstruction `json:"instructions"`
}

type ParsedMessage struct {
	AccountKeys     []ParsedMessageAccount `json:"accountKeys"`
	Instructions    []ParsedInstruction    `json:"instructions"`
	RecentBlockhash string                 `json:"recentBlockhash"`
}

type ParsedTransaction struct {
	Message    ParsedMessage `json:"message"`
	Signatures []string      `json:"signatures"`
}

type ParsedTransactionMeta struct {
	// Error if transaction failed, null if transaction succeeded.
	// https://github.com/solana-labs/solana/blob/master/sdk/src/transaction.rs#L24
	Err interface{} `json:"err"`

	// Fee this transaction was charged
	Fee uint64 `json:"fee"`

	// Array of u64 account balances from before the transaction was processed
	PreBalances []uint64 `json:"preBalances"`

	// Array of u64 account balances after the transaction was processed
	PostBalances []uint64 `json:"postBalances"`

	// List of inner instructions or omitted if inner instruction recording
	// was not yet enabled during this transaction
	InnerInstructions []ParsedInnerInstruction `json:"innerInstructions"`

	// List of token balances from before the transaction was processed
	// or omitted if token balance recording was not yet enabled during this transaction
	PreTokenBalances []TokenBalance `json:"preTokenBalances"`

	// List of token balances from after the transaction was processed
	// or omitted if token balance recording was not yet enabled during this transaction
	PostTokenBalances []TokenBalance `json:"postTokenBalances"`

	// Array of string log messages or omitted if log message
	// recording was not yet enabled during this transaction
	LogMessages []string `json:"logMessages"`
}

type ParsedTransactionWithMeta struct {
	Slot        uint64                `json:"slot"`
	BlockTime   int64                 `json:"blockTime"`
	Transaction ParsedTransaction     `json:"transaction"`
	Meta        ParsedTransactionMeta `json:"meta"`
}

type RewardType string

const (
	RewardTypeFee     RewardType = "Fee"
	RewardTypeRent    RewardType = "Rent"
	RewardTypeVoting  RewardType = "Voting"
	RewardTypeStaking RewardType = "Staking"
)

type BlockReward struct {
	// The public key of the account that received the reward.
	Pubkey string `json:"pubkey"`

	// Number of reward lamports credited or debited by the account, as a i64.
	Lamports int64 `json:"lamports"`

	// Account balance in lamports after the reward was applied.
	PostBalance uint64 `json:"postBalance"`

	// Type of reward: "Fee", "Rent", "Voting", "Staking".
	RewardType RewardType `json:"rewardType"`

	// Vote account commission when the reward was credited,
	// only present for voting and staking rewards.
	Commission *uint8 `json:"commission,omitempty"`
}

