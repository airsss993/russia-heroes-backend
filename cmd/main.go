package main

import (
	"fmt"

	"github.com/airsss993/russia-heroes-backend/pkg/utils"
)

func main() {
	secret, _ := utils.GenerateSecretKey()
	id, _ := utils.GenerateAccessKeyID()
	fmt.Println(secret, id)
}
