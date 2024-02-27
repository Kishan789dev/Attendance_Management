package login

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg"

	"github.com/kk/attendance_management/authentication/getrole"
	bean "github.com/kk/attendance_management/bean"
)

type LoginRest interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type LoginRestImpl struct {
	loginsvc LoginSvc
	getrole  getrole.GetroleRest
}

func NewLoginRest(loginsvc LoginSvc, getrole getrole.GetroleRest) *LoginRestImpl {
	return &LoginRestImpl{
		loginsvc: loginsvc,
		getrole:  getrole,
	}
}

var jwtKey = []byte("secret_key")

func (impl *LoginRestImpl) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("kxfkjdfhjkdfghkjdfhkj")
	var credentialStudent bean.Credentials

	err := json.NewDecoder(r.Body).Decode(&credentialStudent)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err, mypass := impl.loginsvc.LoginSvc(credentialStudent.Useremail)
	if err != nil {
		if err == pg.ErrNoRows {
			errMsg := make(map[string]string, 0)
			errMsg["error"] = "wrong details entered"

			json.NewEncoder(w).Encode(errMsg)
			return
		}
		fmt.Println("error found", err)
		errMsg := make(map[string]string, 0)
		errMsg["error"] = "error due to server search"

		json.NewEncoder(w).Encode(errMsg)
		return
	}
	fmt.Println("Expected password")
	expectedPassword := credentialStudent.Password
	actualPassword := mypass

	if expectedPassword != actualPassword {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	expirationTime := time.Now().Add(time.Hour * 999999)

	claims := &bean.Claims{
		Useremail: credentialStudent.Useremail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:     "token",
			Value:    tokenString,
			Expires:  expirationTime,
			HttpOnly: false,
			Secure:   false,
			Domain:   "",
			Path:     "/",
		})

	role, err := impl.getrole.GetRoletemp(w, r, tokenString)
	if err == nil {
		errorMap := map[string]int{
			"role": role,
		}

		json.NewEncoder(w).Encode(errorMap)
		return

	}
}
