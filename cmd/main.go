package main

import (
	"fmt"

	"github.com/AlGrushino/subscribes/pkg/db"
	"github.com/AlGrushino/subscribes/pkg/handlers"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg, err := db.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	db, err := db.DBInit(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	userID, err := uuid.Parse("c02dadb7-8f30-451f-ad81-ea17f127d33b")
	if err != nil {
		fmt.Println(err)
		return
		// log.Fatal("Invalid UUID:", err)
	}
	subscribes, err := handlers.GetUserSubscribes(userID, db)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Найдено %d подписок для пользователя %s:\n", len(subscribes), userID)
	for _, sub := range subscribes {
		endDateStr := "бессрочная"
		if sub.EndDate != nil {
			endDateStr = sub.EndDate.Format("2006-01-02")
		}
		fmt.Printf("- %s: %d руб. (с %s по %s)\n",
			sub.ServiceName,
			sub.Price,
			sub.StartDate.Format("2006-01-02"),
			endDateStr)
	}
}
