package teachers

import (
	"encoding/json"
	"net/http"

	"github.com/go-pg/pg"
	"github.com/kk/attendance_management/authentication/getrole"
	token "github.com/kk/attendance_management/authentication/tokenvalidation"
	"github.com/kk/attendance_management/components/users"

	bean "github.com/kk/attendance_management/bean"
)

type TeacherRest interface {
	AddTeacher(w http.ResponseWriter, r *http.Request)
	TeacherEntryPunchin(w http.ResponseWriter, r *http.Request)
	TeacherEntryPunchOut(w http.ResponseWriter, r *http.Request)
	GetTeacherattendance(w http.ResponseWriter, r *http.Request)
	GetClassattendance(w http.ResponseWriter, r *http.Request)
}

type TeacherRestImpl struct {
	teachersvc TeacherSvc
	getrole    getrole.GetroleRest
	user       users.UserRest
	valid      token.AuthenticationRest
}

func NewteacherRest(teachersvc TeacherSvc, getrole getrole.GetroleRest, user users.UserRest, valid token.AuthenticationRest) *TeacherRestImpl {
	return &TeacherRestImpl{
		teachersvc: teachersvc,
		getrole:    getrole,
		user:       user,
		valid:      valid,
	}
}

func (impl *TeacherRestImpl) AddTeacher(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	role, err := impl.getrole.GetRole(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if role != 3 {
		json.NewEncoder(w).Encode("only principle can add teacher")
		return
	}

	var userdetails bean.Userdetails

	err = json.NewDecoder(r.Body).Decode(&userdetails)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	teacher := bean.Teacher{Name: userdetails.Name, Address: userdetails.Address, Email: userdetails.Email}
	err = impl.teachersvc.AddTeacherSvc(&teacher, &userdetails)

	if err != nil {

		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode("teacher added successfully")

}

func (impl *TeacherRestImpl) TeacherEntryPunchin(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	email, err := impl.valid.ValidateTokenAndGetEmail(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err, CustomErrTyp, tid := impl.teachersvc.TeacherEntryPunchinSvc(email)

	if CustomErrTyp == 0 {
		json.NewEncoder(w).Encode("you are not a Teacher")
		return
	}

	var aid int
	if CustomErrTyp == 1 {

		err, aid = impl.teachersvc.TeacherAttendanceWithPunchData(tid)

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
	aid = tid
	err, str := impl.teachersvc.TeacherPunchEntryInTable(aid)

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

func (impl *TeacherRestImpl) TeacherEntryPunchOut(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	email, err := impl.valid.ValidateTokenAndGetEmail(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err, CustomErrTyp, _, aid := impl.teachersvc.TeacherEntryPunchOutSvc(email)
	if CustomErrTyp == 0 {
		json.NewEncoder(w).Encode("you are not a Teacher")
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

	err, str := impl.teachersvc.TeacherPunchOutEntryInTable(aid)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(str)

}

func (impl *TeacherRestImpl) GetTeacherattendance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role, err := impl.getrole.GetRole(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	email, err := impl.valid.ValidateTokenAndGetEmail(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var tid int

	if role == 1 {
		json.NewEncoder(w).Encode("u are a student so u cant get details")
		return

	}
	if role == 2 {

		err = impl.teachersvc.GetTeacherattendanceSvcTidGetting(email, &tid)

		if err != nil {
			role = 1

			json.NewEncoder(w).Encode(err.Error())
			return
		}

	}

	if role == 2 || role == 3 {
		teacherattendance := &bean.TeacherAttendance{}

		json.NewDecoder(r.Body).Decode(&teacherattendance)

		if teacherattendance.Tid != 0 {
			tid = teacherattendance.Tid
		}

		err, teacherattendancedetail := impl.teachersvc.GetTeacherAttendanceDetailsSvc(tid, teacherattendance.Month, teacherattendance.Year)

		if err != nil {

			errMsg := make(map[string]string, 0)
			errMsg["error"] = err.Error()

			json.NewEncoder(w).Encode(errMsg)
			return

		} else {

			json.NewEncoder(w).Encode(teacherattendancedetail)
			return
		}
	} else {
		json.NewEncoder(w).Encode("you are not a teacher")
		return
	}

}

func (impl *TeacherRestImpl) GetClassattendance(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	role, err := impl.getrole.GetRole(w, r)
	if err != nil {
		json.NewEncoder(w).Encode("you are not a teacher")

		return
	}
	if role == 1 || role == 3 {
		json.NewEncoder(w).Encode("you are not a teacher")
		return

	}

	var classtemp = bean.Classtemp{}

	json.NewDecoder(r.Body).Decode(&classtemp)

	err, classdata := impl.teachersvc.GetClassattendanceSvc(&classtemp)

	if err != nil {
		errMsg := make(map[string]string, 0)
		errMsg["error"] = err.Error()

		json.NewEncoder(w).Encode(errMsg)
		return

	} else {
		if classdata != nil {
			json.NewEncoder(w).Encode(classdata)

		} else {
			errorMap := map[string]string{
				"error": "invalid inputs",
			}
			json.NewEncoder(w).Encode(errorMap)

		}
		return
	}

}
