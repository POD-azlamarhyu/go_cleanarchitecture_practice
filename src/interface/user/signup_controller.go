package user

import (
	"cleanarchitecture-practice/src/application/user"
	"encoding/json"
	"net/http"
	"cleanarchitecture-practice/src/interface/shared"
)

type SignupController struct {
	signupUseCase *user.ISignUpUseCase
}

func NewSignupController(signupUseCase user.ISignUpUseCase) *SignupController {
	return &SignupController{
		signupUseCase: &signupUseCase,
	}
}

func (c *SignupController) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var input SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		shared.RespondWithBadRequest(rw, "Invalid request body")
		return
	}

	output, err := (*c.signupUseCase).Execute(r.Context(), RequestToInputDTO(&input))
	if err != nil {
		shared.RespondWithInternalServerError(rw, err.Error())
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(OutputDTOToResponse(output))
}