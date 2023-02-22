package pyl

type CheckBalance struct {
	UserId string `json:"userId"`
}

type AddTransaction struct {
	UserId    string  `json:"userId"`
	Amount    float64 `json:"amount"`
	Flow      string  `json:"flow"`
	TypeTrans string  `json:"typeTrans"`
	CreateBy  string  `json:"createBy"`
	ApproveBy string  `json:"approveBy"`
	Desc      string  `json:"desc"`
}
