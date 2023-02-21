package pyl

type CreateMember struct {
	UserId   string `json:"userId"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
