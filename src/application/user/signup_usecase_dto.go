package user

type SignUpInputDTO struct{
	Username string
	Email    string
	Password string
}

type SignUpOutputDTO struct{
	UserId string
	Username string
	Email string
}