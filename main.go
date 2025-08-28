package main

import (
	"fmt"

	"github.com/lukaspodobnik/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg.SetUser("Lukas")

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cfg.DataBankURL)
	fmt.Println(cfg.CurrentUserName)
}
