package tokenvalid

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/kk/attendance_management/bean"
)

var jwtKey = []byte("secret_key")

func ValidateTokenAndGetEmail(w http.ResponseWriter, r *http.Request) (string, error) {
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
	// check email expiration

	// var newuser bean.User
	// db := dataBase.Connect()
	// defer db.Close()
	// err = db.Model(&newuser).Where("email=?", claims.Useremail).Select()
	// if err == pg.ErrNoRows {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return "", err
	// }
	// if bean.Claims.ExpiresAt
	err = ValidateTokenAndGetEmailSvc(claims.Useremail)
	if err != nil {
		return "", err
	}

	return claims.Useremail, err

}

func ValidateTokenAndGetEmailtemp(w http.ResponseWriter, r *http.Request, tokenstr string) (string, error) {
	// cookie, err := r.Cookie("token")
	// fmt.Println("cookie is ", cookie)
	// if err != nil {
	// 	if err == http.ErrNoCookie {
	// 		return "", err
	// 	}
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return "", err
	// }

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
	// check email expiration

	err = ValidateTokenAndGetEmailSvc(claims.Useremail)
	if err != nil {
		return "", err
	}

	return claims.Useremail, err
	// if bean.Claims.ExpiresAt

}
