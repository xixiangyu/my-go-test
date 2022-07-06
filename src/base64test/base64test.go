package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "ac1234"
	s2 := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(s2)

	b, _ := base64.StdEncoding.DecodeString("WVZsMlUxRVE4ZA==")
	fmt.Println(string(b))
}
