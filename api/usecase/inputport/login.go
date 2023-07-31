package inputport

import "echo-admin-templte/entity"

type Login interface {
	Login(e entity.Login) (entity.LoginToken, error)
}
