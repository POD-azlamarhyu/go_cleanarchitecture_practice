package server

import (
	userUsecase "cleanarchitecture-practice/src/application/user"
	"cleanarchitecture-practice/src/infrastructure/database"
	userInfra "cleanarchitecture-practice/src/infrastructure/user"
	"cleanarchitecture-practice/src/interface/user"
)

var (
	signupController *user.SignupController
)

func initControllers() {
	initUserControllers()
}

func initUserControllers() {
	signupController = user.NewSignupController(
		userUsecase.NewSignUpUseCase(
			userInfra.NewUserRepository(database.GetDammyDB()),
		),
	)
}
