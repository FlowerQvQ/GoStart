package main

import (
	_ "github.com/google/wire"
)

func main() {
	engin := InitApp()

	err := engin.Run()
	if err != nil {
		return
	}

}
