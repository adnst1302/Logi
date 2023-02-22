package pyl

type CreateMember struct {
	UserId   string `json:"userId"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type UpdateProfileMember struct {
	UserId         string `json:"userId"`
	FullName       string `json:"fullName"`
	PlaceOfBirth   string `json:"placeOfBirth"`
	BirthDate      string `json:"birthDate"`
	PhoneNumber    string `json:"phoneNumber"`
	Address        string `json:"address"`
	EducationLevel string `json:"educationLevel"`
	Division       string `json:"division"`
	SignInDate     string `json:"signInDate"`
	StatusEmployee string `json:"statusEmployee"`
}
