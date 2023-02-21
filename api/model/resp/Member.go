package resp

type CreateMember struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

type SuccessCreateMember struct {
	Message string `json:"message"`
}

type GetAllMember struct {
	Success bool        `json:"status"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}
