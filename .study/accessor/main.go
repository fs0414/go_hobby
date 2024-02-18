package main

import (
	"fmt"

	"github.com/fs0414/go_hobby/.study/accessor/user"
)

func main() {
	result := get.User{Name: "user01"}
	fmt.Println(result.Name)
}

