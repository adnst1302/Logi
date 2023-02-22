package resp

type CheckBalance struct {
	Success bool    `json:"success"`
	Code    int     `json:"code"`
	Balance float64 `json:"balance"`
	Message string  `json:"message"`
}
