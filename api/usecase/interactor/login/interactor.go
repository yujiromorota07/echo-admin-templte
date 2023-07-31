package login

import (
	"context"
	"echo-admin-templte/entity"
	constant "echo-admin-templte/entity/constants"
	"echo-admin-templte/usecase/inputport"
	"echo-admin-templte/usecase/repository"
	util "echo-admin-templte/utils"
	"fmt"
)

type LoginUsecase struct {
	accountRepository repository.AccountRepository
}

func NewLoginUsecase(accountRepository repository.AccountRepository) inputport.Login {
	return &LoginUsecase{
		accountRepository: accountRepository,
	}
}

func (usecase LoginUsecase) Login(e entity.Login) (entity.LoginToken, error) {
	ctx := context.Background()
	//アカウント確認
	account, err := usecase.accountRepository.GetAccountByLogin(ctx, e)
	if err != nil {
		return entity.LoginToken{}, entity.NewBadRequestError(constant.LoginBadRequestMessage)
	}

	// //パスワード確認
	password := util.Password{
		Password: e.Password,
		Hash:     account.Password,
	}

	err = password.CompareHashAndPassword()
	if err != nil {
		fmt.Println(err)
		return entity.LoginToken{}, entity.NewBadRequestError(constant.LoginBadRequestMessage)
	}

	//トークン作成
	loginClaim := util.LoginClaim{
		AccountID:   account.ID,
		AccountName: account.Name,
		Email:       account.Email,
	}

	token, err := util.GenerateToken(loginClaim)
	if err != nil {
		return entity.LoginToken{}, err
	}

	return entity.LoginToken{Token: token}, nil
}
