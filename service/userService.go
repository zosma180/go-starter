package service

import (
	"fmt"

	"github.com/zosma180/go-poc/dto"
)

func CreateUser(model dto.CreateUserRequest) error {
	fmt.Println("Fake User Created")
	return nil
}

func GetUserList() []string {
	list := []string{"A", "B", "C"}
	return list
}
