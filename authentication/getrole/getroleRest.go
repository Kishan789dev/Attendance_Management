package getrole

import (
	"encoding/json"
	"net/http"

	token "github.com/kk/attendance_management/authentication/tokenvalidation"
	// login "github.com/kk/attendance_management/authentication/login"
)

func GetRole(w http.ResponseWriter, r *http.Request) (int, error) {

	email, err := token.ValidateTokenAndGetEmail(w, r)
	if err != nil {
		json.NewEncoder(w).Encode("user is unauthorised")
		return 0, err

	}

	err, role := GetRolesvc(email)

	if err != nil {

		errMsg := make(map[string]string, 0)
		errMsg["error"] = err.Error()
		json.NewEncoder(w).Encode(errMsg)
		return 0, err
	}
	return role, err

}

func GetRoletemp(w http.ResponseWriter, r *http.Request, tokenstr string) (int, error) {

	email, err := token.ValidateTokenAndGetEmailtemp(w, r, tokenstr)
	if err != nil {
		json.NewEncoder(w).Encode("user is unauthorised")
		return 0, err

	}
	err, role := GetRolesvc(email)

	if err != nil {

		errMsg := make(map[string]string, 0)
		errMsg["error"] = err.Error()
		json.NewEncoder(w).Encode(errMsg)
		return 0, err
	}
	return role, err

}
