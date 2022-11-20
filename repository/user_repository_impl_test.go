package repository

import (
	"final-project-backend/config"
	"final-project-backend/entity"
	"final-project-backend/pkg/database/postgres"
	"fmt"
	"testing"
)

var dbcfg = config.DatabaseConfig{
	Host:     "localhost",
	Port:     "5432",
	DbName:   "final_project",
	User:     "dedyirawan",
	Password: "dedyirawan",
}
var db = postgres.New(&config.Config{Database: dbcfg})

func TestUserRepositoryImpl_UpdateUser(t *testing.T) {
	user := entity.User{
		Id:       2,
		Email:    "dedy",
		FullName: "dedy",
		Phone:    "1235",
		Photo:    "tespu",
	}

	r := NewUserRepositoryImpl(db)
	res, err := r.UpdateUser(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Email)
}

func TestUserRepositoryImpl_AddBalance(t *testing.T) {
	r := NewUserRepositoryImpl(db)
	res, err := r.AddBalance(1, 50000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Balance)
}

func TestUserRepositoryImpl_ReduceBalance(t *testing.T) {
	r := NewUserRepositoryImpl(db)
	res, err := r.ReduceBalance(1, 20000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Balance)
}
