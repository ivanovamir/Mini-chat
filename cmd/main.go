package main

import (
	"fmt"
	"mini_chat/internal/configs"
)

func main() {
	//TODO: implement config +
	//TODO: implement logger (slog)
	//TODO: implement handler
	//TODO: implement server
	//TODO: implement db

	cfg := configs.NewConfig("configs/config.yml")

	fmt.Println(cfg)
}
