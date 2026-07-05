package user


import (
	"cleanarchitecture-practice/src/domain/user"
	"context"
)

type ISignUpUseCase interface {
	Execute(ctx context.Context, input *SignUpInputDTO) (*SignUpOutputDTO, error)
}


type SignUpUseCase struct {
	userRepository user.IUserRepository
}

func NewSignUpUseCase(userRepository user.IUserRepository) ISignUpUseCase {
	return &SignUpUseCase{
		userRepository: userRepository,
	}
}

func (uc *SignUpUseCase) Execute(ctx context.Context, input *SignUpInputDTO) (*SignUpOutputDTO, error){
	user, err := user.NewUser(
		input.Username,
		input.Email,
		input.Password,
	)
	if err != nil {
		return nil, err
	}

	err = uc.userRepository.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	return &SignUpOutputDTO{
		UserId:   user.GetUserId().String(),
		Username: user.GetUserName(),
		Email:    user.GetEmail(),
	}, nil
}