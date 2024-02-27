package tokenvalid

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/kk/attendance_management/bean"
)

type AuthenticationRest interface {
	ValidateTokenAndGetEmail(w http.ResponseWriter, r *http.Request) (string, error)
	ValidateTokenAndGetEmailtemp(w http.ResponseWriter, r *http.Request, tokenstr string) (string, error)
}
type AuthenticationRestImpl struct {
	authenticationsvc AuthenticationSvc
}

func NewAuthenticationRest(authenticationsvc AuthenticationSvc) *AuthenticationRestImpl {
	return &AuthenticationRestImpl{
		authenticationsvc: authenticationsvc,
	}
}

var jwtKey = []byte("secret_key")

func (impl *AuthenticationRestImpl) ValidateTokenAndGetEmail(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie("token")
	fmt.Println("cookie is ", cookie)
	if err != nil {
		if err == http.ErrNoCookie {
			return "", err
		}
		w.WriteHeader(http.StatusBadRequest)
		return "", err
	}

	tokenStr := cookie.Value

	claims := &bean.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return "", err
		}
		w.WriteHeader(http.StatusBadRequest)
		return "", err
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return "", err
	}

	err = impl.authenticationsvc.ValidateTokenAndGetEmailSvc(claims.Useremail)
	if err != nil {
		return "", err
	}

	return claims.Useremail, err

}

func (impl *AuthenticationRestImpl) ValidateTokenAndGetEmailtemp(w http.ResponseWriter, r *http.Request, tokenstr string) (string, error) {

	tokenStr := tokenstr

	claims := &bean.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return "", err
		}
		w.WriteHeader(http.StatusBadRequest)
		return "", err
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return "", err
	}

	err = impl.authenticationsvc.ValidateTokenAndGetEmailSvc(claims.Useremail)
	if err != nil {
		return "", err
	}

	return claims.Useremail, err

}
