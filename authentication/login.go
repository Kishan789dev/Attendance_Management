package authentication

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg"
	bean "github.com/kk/attendance_management/bean"

	"github.com/kk/attendance_management/dataBase"
)

var jwtKey = []byte("secret_key")

func Login(w http.ResponseWriter, r *http.Request) {
	// log.Println("jsdjfklsdjflsdjf")

	var credentialStudent bean.Credentials
	// var credentialTeacher bean.Teacher
	err := json.NewDecoder(r.Body).Decode(&credentialStudent)
	// fmt.Println(credentialStudent)
	// err = json.NewDecoder(r.Body).Decode(&credentialTeacher)
	var getpassword bean.User

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	db := dataBase.Connect()

	defer db.Close()
	var mypass string

	err = db.Model(&getpassword).Column("password").Where("email=?", credentialStudent.Useremail).Select(&mypass)
	if err != nil {
		// log.Println("k0")
		if err == pg.ErrNoRows {
			fmt.Println("line 38 status u")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	expectedPassword := credentialStudent.Password
	actualPassword := mypass
	// log.Println("k1")

	if expectedPassword != actualPassword {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// log.Println("k2")

	// expectedPassword, ok := users[credentials.Username]

	expirationTime := time.Now().Add(time.Hour * 1000000)

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
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
	// log.Println("kisahn")

}

func ValidateTokenAndGetEmail(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie("token")

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
