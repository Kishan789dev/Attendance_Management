package students

import (
	"encoding/json"

	"strings"

	"net/http"

	"github.com/go-pg/pg"
	token "github.com/kk/attendance_management/authentication/tokenvalidation"

	"github.com/kk/attendance_management/authentication/getrole"
	"github.com/kk/attendance_management/components/users"

	bean "github.com/kk/attendance_management/bean"
)

type StudentRest interface {
	AddStudent(w http.ResponseWriter, r *http.Request)
	StudentEntryPunchin(w http.ResponseWriter, r *http.Request)
	StudentEntryPunchOut(w http.ResponseWriter, r *http.Request)
	GetStudentattendance(w http.ResponseWriter, r *http.Request)
}

type StudentRestImpl struct {
	studentsvc StudentService
	getrole    getrole.GetroleRest
	user       users.UserRest
	valid      token.AuthenticationRest
}

func NewStudentRest(data StudentService, getrole getrole.GetroleRest, user users.UserRest, valid token.AuthenticationRest) *StudentRestImpl {
	return &StudentRestImpl{
		studentsvc: data,
		getrole:    getrole,
		user:       user,
		valid:      valid,
	}
}

func ExtractStaticPart(errorMessage string) string {
	index := strings.Index(errorMessage, ":")
	if index != -1 {
		return strings.TrimSpace(errorMessage[:index])
	}
	return errorMessage
}

func (impl *StudentRestImpl) AddStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	role, err := impl.getrole.GetRole(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if role != 3 {
		json.NewEncoder(w).Encode("only principle can add student")
		return
	}

	var userdetails bean.Userdetails
	err = json.NewDecoder(r.Body).Decode(&userdetails)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	student := bean.Student{
		Name:    userdetails.Name,
		Address: userdetails.Address,
		Class:   userdetails.Class,
		Email:   userdetails.Email,
	}

	err = impl.studentsvc.AddStudentService(&student, &userdetails)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = impl.user.AddUser(userdetails.Email, 1, userdetails.Password)
	if err != nil {

		json.NewEncoder(w).Encode("error in adding user to user table")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(" Student added suceesfully")
}

func (impl *StudentRestImpl) StudentEntryPunchin(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	email, err := impl.valid.ValidateTokenAndGetEmail(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err, CustomErrTyp, sid := impl.studentsvc.StudentEntryPunchinSvc(email)

	if CustomErrTyp == 0 {
		json.NewEncoder(w).Encode("you are not a student")
		return
	}

	var aid int
	if CustomErrTyp == 1 {

		err, aid = impl.studentsvc.StudentAttendanceWithPunchData(sid)

		if err != nil {

			w.WriteHeader(http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode("Success")
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	aid = sid
	err, str := impl.studentsvc.StudentPunchEntryInTable(aid)

	if err != nil {
		if len(str) > 0 {
			json.NewEncoder((w)).Encode(str)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}
	json.NewEncoder(w).Encode(str)
}

func (impl *StudentRestImpl) StudentEntryPunchOut(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	email, err := impl.valid.ValidateTokenAndGetEmail(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err, CustomErrTyp, _, aid := impl.studentsvc.StudentEntryPunchOutSvc(email)
	if CustomErrTyp == 0 {
		json.NewEncoder(w).Encode("you are not a student")
		return
	}

	if err != nil && err != pg.ErrNoRows {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err == pg.ErrNoRows && CustomErrTyp == 1 {

		json.NewEncoder(w).Encode(" no data found  so go for punch in first")

		return

	}

	err, str := impl.studentsvc.StudentPunchOutEntryInTable(aid)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(str)

}

func (impl *StudentRestImpl) GetStudentattendance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	email, err := impl.valid.ValidateTokenAndGetEmail(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("kindly login")
		return
	}

	err, sid := impl.studentsvc.GetStudentsAttendanceEmailSvc(email)

	if err != nil {
		json.NewEncoder(w).Encode("you are not a student")
		return
	}

	studentattendance := &bean.StudentAttendance{Sid: sid}
	json.NewDecoder(r.Body).Decode(&studentattendance)

	studentattendancedetail, err := impl.studentsvc.FetchAttendanceFromDetailsSvc(studentattendance)

	if err != nil {
		staticPart := ExtractStaticPart(err.Error())

		json.NewEncoder(w).Encode(staticPart)

		w.WriteHeader(http.StatusBadRequest)
		return

	} else {

		if len(*studentattendancedetail) == 0 {

			mapp := map[string]string{
				"error": "student with this details doesn't exist",
			}
			json.NewEncoder(w).Encode(mapp)
			return
		}

		json.NewEncoder(w).Encode(studentattendancedetail)
		return
	}

}
