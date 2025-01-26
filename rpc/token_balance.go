package rpc

type UiTokenAmount struct {
	// Raw amount of tokens as a string, ignoring decimals.
	Amount string `json:"amount"`

	// Number of decimals configured for token's mint.
	Decimals int64 `json:"decimals"`

	// DEPRECATED: Token amount as a float, accounting for decimals.
	UiAmount *float64 `json:"uiAmount"`

	// Token amount as a string, accounting for decimals.
	UiAmountString string `json:"uiAmountString"`
}

type TokenBalance struct {
	// Index of the account in which the token balance is provided for.
	AccountIndex uint16 `json:"accountIndex"`

	// Pubkey of token balance's owner.
	Owner string `json:"owner,omitempty"`
	// Pubkey of token program.
	ProgramId string `json:"programId,omitempty"`

	// Pubkey of the token's mint.
	Mint string `json:"mint"`

	UiTokenAmount *UiTokenAmount `json:"uiTokenAmount"`
}
