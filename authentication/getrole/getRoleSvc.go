package getrole

func GetRolesvc(email string) (error, int) {
	return GetRoleRepo(email)
}
