package main

import (
	"fmt"
	"strings"
)

func main (){
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string{
fields := strings.Fields(text)
	return fields
}