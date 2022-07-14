package main

import (
	"encoding/base64"
	"fmt"
)

func reversebyte(input []byte) []byte {
	if len(input) == 0 {
		return input
	}
	return append(reversebyte(input[1:]), input[0])
}

func main() {
	var whatIsIt string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)

	whatIsIt = string(reversebyte(sd))
	fmt.Println(whatIsIt)
}
