package rpc

const SOLANA_MAINNET_RPC_URL = "https://api.mainnet-beta.solana.com"

type Commitment string

const (
	CommitmentFinalized Commitment = "finalized"
	CommitmentConfirmed Commitment = "confirmed"
	CommitmentProcessed Commitment = "processed"
)

type Encoding string

const (
	EncodingBase58     Encoding = "base58"
	EncodingBase64     Encoding = "base64"
	EncodingBase64Zstd Encoding = "base64+zstd"
	EncodingJson       Encoding = "json"
	EncodingJsonParsed Encoding = "jsonParsed"
)

type TransactionVersion int64

const (
	TransactionVersionV0 TransactionVersion = 0

	maxSupportedTransactionVersion TransactionVersion = TransactionVersionV0
)

type TransactionDetails string

const (
	TransactionDetailsFull       TransactionDetails = "full"
	TransactionDetailsAccounts   TransactionDetails = "accounts"
	TransactionDetailsSignatures TransactionDetails = "signatures"
	TransactionDetailsNone       TransactionDetails = "none"
)
