package user

import (
	"cleanarchitecture-practice/src/application/user"
)

type SignUpRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func RequestToInputDTO(req *SignUpRequest) *user.SignUpInputDTO {
	return &user.SignUpInputDTO{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}

func OutputDTOToResponse(output *user.SignUpOutputDTO) *SignUpResponse {
	return &SignUpResponse{
		UserId:   output.UserId,
		Username: output.Username,
		Email:    output.Email,
	}
}

