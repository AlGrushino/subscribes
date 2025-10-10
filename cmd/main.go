package main

import (
	"fmt"

	"github.com/AlGrushino/subscribes/pkg/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
		return
	}

	users, err := handlers.GetUserByAge(21)
	if err != nil {
		problem := fmt.Errorf("ошибка запроса: %v", err)
		fmt.Println(problem)
		return
	}

	for _, v := range users {
		fmt.Printf("имя: %s\nвозраст: %d\n", v.Name, v.Age)
	}
}
