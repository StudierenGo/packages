package main

import (
	"demo/packages/account"
	"demo/packages/helpers"
	"fmt"
)

func main() {
	userLogin, userPassword, userUrl := helpers.GetUserInput()
	account, err := account.UserAccountWithTimeStampConstructor(userLogin, userPassword, userUrl)

	if err != nil {
		fmt.Println("Error creating account:", err)
		return
	}

	result := account.OutputUserData()
	fmt.Println(result)
}
