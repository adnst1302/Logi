package resp

type CreateMember struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

// CreateMember.DataCreateMember
type DataCreateMember struct {
	Message string `json:"message"`
}

type GetAllMember struct {
	Success bool               `json:"status"`
	Code    int                `json:"code"`
	Data    []DataGetAllMember `json:"data"`
	Message string             `json:"message"`
}

// GetAllMember.DataGetAllMember
type DataGetAllMember struct {
	UserId  string                  `json:"userId"`
	Email   string                  `json:"email"`
	Role    string                  `json:"role"`
	Profile ProfileDataGetAllMember `json:"profile"`
}

type ProfileDataGetAllMember struct {
	FullName     string `json:"fullName"`
	PlaceOfBirth string `json:"placeOfBirth"`
	BirthDate    string `json:"birthDate"`
	PhoneNumber  string `json:"phoneNumber"`
	Address      string `json:"address"`
}
