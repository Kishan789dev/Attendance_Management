//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"github.com/kk/attendance_management/authentication/getrole"
	"github.com/kk/attendance_management/authentication/login"
	"github.com/kk/attendance_management/routers"

	token "github.com/kk/attendance_management/authentication/tokenvalidation"
	"github.com/kk/attendance_management/components/students"
	"github.com/kk/attendance_management/components/teachers"
	"github.com/kk/attendance_management/components/users"
	"github.com/kk/attendance_management/dataBase"
)

func InitialiseApp() *App {
	wire.Build(
		// database
		mux.NewRouter,

		dataBase.NewDataBase,
		wire.Bind(new(dataBase.DataBase), new(*dataBase.DataBaseImpl)),

		getrole.NewGetroleRest,
		wire.Bind(new(getrole.GetroleRest), new(*getrole.GetroleRestImpl)),
		getrole.NewGetroleSvc,
		wire.Bind(new(getrole.GetroleSvc), new(*getrole.GetroleSvcImpl)),
		getrole.NewGetRole,
		wire.Bind(new(getrole.GetroleRepo), new(*getrole.GetroleRepoImpl)),

		// login
		login.NewLoginRest,
		wire.Bind(new(login.LoginRest), new(*login.LoginRestImpl)),
		login.NewLoginSvc,
		wire.Bind(new(login.LoginSvc), new(*login.LoginSvcImpl)),
		login.NewLoginRepo,
		wire.Bind(new(login.LoginRepo), new(*login.LoginRepoImpl)),

		//authentication
		token.NewAuthenticationRest,
		wire.Bind(new(token.AuthenticationRest), new(*token.AuthenticationRestImpl)),
		token.NewAuthenticationSvc,
		wire.Bind(new(token.AuthenticationSvc), new(*token.AuthenticationSvcImpl)),
		token.NewAuthenticationRepo,
		wire.Bind(new(token.AuthenticationRepo), new(*token.AuthenticationRepoImpl)),

		//students
		students.NewStudentRest,
		wire.Bind(new(students.StudentRest), new(*students.StudentRestImpl)),

		students.NewStudentservice,
		wire.Bind(new(students.StudentService), new(*students.StudentServiceImpl)),

		students.NewStudentRepository,
		wire.Bind(new(students.StudentRepository), new(*students.StudentRepositoryImpl)),

		// Teachers
		teachers.NewteacherRest,
		wire.Bind(new(teachers.TeacherRest), new(*teachers.TeacherRestImpl)),
		teachers.NewTeacherSvc,
		wire.Bind(new(teachers.TeacherSvc), new(*teachers.TeacherSvcImpl)),
		teachers.NewTeacherRepositoryImpl,
		wire.Bind(new(teachers.TeacherRepo), new(*teachers.TeacherRepositoryImpl)),

		//users
		users.NewUserRestImpl,
		wire.Bind(new(users.UserRest), new(*users.UserRestImpl)),
		users.NewUserSvc,
		wire.Bind(new(users.UserSvc), new(*users.UserSvcImpl)),
		users.NewUserRepo,
		wire.Bind(new(users.UserRepo), new(*users.UserRepoImpl)),
		routers.NewRoute,
		NewApp,
	)
	return &App{}

}
