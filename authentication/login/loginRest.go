package login

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg"
	bean "github.com/kk/attendance_management/bean"

	"github.com/kk/attendance_management/authentication"
	"github.com/kk/attendance_management/dataBase"
)

var jwtKey = []byte("secret_key")

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var credentialStudent bean.Credentials

	err := json.NewDecoder(r.Body).Decode(&credentialStudent)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// var getpassword bean.User

	// err = db.Model(&getpassword).Column("password").Where("email=?", credentialStudent.Useremail).Select(&mypass)

	err, mypass := LoginSvc(credentialStudent.Useremail)
	if err != nil {
		// log.Println("k0")
		if err == pg.ErrNoRows {
			// fmt.Println("line 38 status u")
			// w.WriteHeader(http.StatusInternalServerError)
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
	// log.Println("k1")

	if expectedPassword != actualPassword {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// log.Println("k2")

	// expectedPassword, ok := users[credentials.Username]

	expirationTime := time.Now().Add(time.Hour * 400000)

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

	// log.Println("kisahn")
	role, err := authentication.GetRoletemp(w, r, tokenString)
	if err == nil {
		errorMap := map[string]int{
			"role": role,
		}

		json.NewEncoder(w).Encode(errorMap)
		return

	}
	// defer db.Close()
}

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

	var newuser bean.User
	db := dataBase.Connect()
	defer db.Close()
	err = db.Model(&newuser).Where("email=?", claims.Useremail).Select()
	if err == pg.ErrNoRows {
		w.WriteHeader(http.StatusUnauthorized)
		return "", err
	}
	// if bean.Claims.ExpiresAt

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

	var newuser bean.User
	db := dataBase.Connect()
	defer db.Close()
	err = db.Model(&newuser).Where("email=?", claims.Useremail).Select()
	if err == pg.ErrNoRows {
		w.WriteHeader(http.StatusUnauthorized)
		return "", err
	}
	// if bean.Claims.ExpiresAt

	return claims.Useremail, err

}
