package main

import (
	"fmt"
)

func main() {
	res, err := MakeQuery()
	if err != nil {
		problem := fmt.Errorf("ошибка запроса: %v", err)
		fmt.Println(problem)
		return
	}

	for _, v := range res {
		fmt.Printf("имя: %s\nвозраст: %d\n", v.Name, v.Age)
	}
}
