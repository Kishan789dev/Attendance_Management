package users

import bn "github.com/kk/attendance_management/bean"

type UserSvc interface {
	AddUserSvc(user *bn.User) error
}
type UserSvcImpl struct {
	userrepo UserRepo
}

func NewUserSvc(userrepo UserRepo) *UserSvcImpl {
	return &UserSvcImpl{
		userrepo: userrepo,
	}
}

func (impl *UserSvcImpl) AddUserSvc(user *bn.User) error {
	return impl.userrepo.AddUserRepo(user)
}
