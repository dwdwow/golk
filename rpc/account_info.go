package rpc

type AccountInfo struct {
	Lamports   uint64 `json:"lamports"`
	Owner      string `json:"owner"`
	State      string `json:"state"`
	Data       string `json:"data"`
	Executable bool   `json:"executable"`
	RentEpoch  uint64 `json:"rentEpoch"`
	Space      uint64 `json:"space"`
}
