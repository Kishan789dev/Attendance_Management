package getrole

import (
	"encoding/json"
	"net/http"

	tokenvalid "github.com/kk/attendance_management/authentication/tokenvalidation"
)

type GetroleRest interface {
	GetRole(w http.ResponseWriter, r *http.Request) (int, error)
	GetRoletemp(w http.ResponseWriter, r *http.Request, tokenstr string) (int, error)
}

type GetroleRestImpl struct {
	getrolesvc GetroleSvc
	auth       tokenvalid.AuthenticationRest
}

func NewGetroleRest(getrolesvc GetroleSvc, auth tokenvalid.AuthenticationRest) *GetroleRestImpl {
	return &GetroleRestImpl{
		getrolesvc: getrolesvc,
		auth:       auth,
	}
}

func (impl *GetroleRestImpl) GetRole(w http.ResponseWriter, r *http.Request) (int, error) {

	email, err := impl.auth.ValidateTokenAndGetEmail(w, r)
	if err != nil {
		json.NewEncoder(w).Encode("user is unauthorised")
		return 0, err

	}

	err, role := impl.getrolesvc.GetRolesvc(email)

	if err != nil {

		errMsg := make(map[string]string, 0)
		errMsg["error"] = err.Error()
		json.NewEncoder(w).Encode(errMsg)
		return 0, err
	}
	return role, err

}

func (impl *GetroleRestImpl) GetRoletemp(w http.ResponseWriter, r *http.Request, tokenstr string) (int, error) {

	email, err := impl.auth.ValidateTokenAndGetEmailtemp(w, r, tokenstr)
	if err != nil {
		json.NewEncoder(w).Encode("user is unauthorised")
		return 0, err

	}
	err, role := impl.getrolesvc.GetRolesvc(email)

	if err != nil {

		errMsg := make(map[string]string, 0)
		errMsg["error"] = err.Error()
		json.NewEncoder(w).Encode(errMsg)
		return 0, err
	}
	return role, err

}
