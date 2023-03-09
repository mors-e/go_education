package main

import (
	"calculator_go/config"
	"fmt"
)

func main() {
	myToken := config.NewBotConfig()
	myDB := config.NewDataBases()

	fmt.Println(myToken.Token, myDB)
}