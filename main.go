package main

import (
	"fmt"

	"github.com/calculi-corp/log"
)

func main() {
	fmt.Println("This is experimental repo, have fun!")
	for {
	}
}

func forTest(a, b int) int {
	r := 10
	r = a + b*r
	if r == 10 {
		return r
	}
	fmt.Println("Testing GX-4884")
	fmt.Println("Testing GX-4884")
	fmt.Println("Testing GX-4884")
	return r
}

type awscreds struct {
	Access   string `json:"access"`
	Secret   string `json:"secret"`
	IAMRole  string `json:"iam_role"`
	Password string `json:"password"`
}

func VStrings() {
	creds := &awscreds{
		Access:   "somekey",
		Secret:   "secret",
		Password: "myawesomepassword",
	}

	log.Infof("creds: %v", creds)
}
