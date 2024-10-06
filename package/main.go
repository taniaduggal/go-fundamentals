package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/taniaduggal/go-fundamentals/auth"
	"github.com/taniaduggal/go-fundamentals/user"
)

func main() {
	auth.LoginWithCredentials("codersgyan", "secret")
	session := auth.GetSession()
	fmt.Println("session", session)

	user := user.User{
		Email: "tan@gmail.com",
		// Name:  "Tan",
	}
	fmt.Println(user.Email, user.Name)

	//agr pacakage ke bahr se hme koi function use krna hai toh hme us package ko capital bnana hai
	color.Blue(user.Email)
}
