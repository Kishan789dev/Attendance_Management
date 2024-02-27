package getrole

type GetroleSvc interface {
	GetRolesvc(email string) (error, int)
}

type GetroleSvcImpl struct {
	getrolerepo GetroleRepo
}

func NewGetroleSvc(getrolerepo GetroleRepo) *GetroleSvcImpl {
	return &GetroleSvcImpl{
		getrolerepo: getrolerepo,
	}
}

func (impl *GetroleSvcImpl) GetRolesvc(email string) (error, int) {
	return impl.getrolerepo.GetRoleRepo(email)
}
