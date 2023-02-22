package resp

type CheckBalance struct {
	Success bool    `json:"success"`
	Code    int     `json:"code"`
	Balance float32 `json:"balance"`
	Message string  `json:"message"`
}

type AddTransaction struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
