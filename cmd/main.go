package main

import (
	"fmt"

	"github.com/AlGrushino/subscribes/pkg/db"
	"github.com/AlGrushino/subscribes/pkg/handlers"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

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

	database, err := db.DBInit(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer database.Close()

	// m, err := migrate.New(
	// 	"file://../migrations",
	// 	db.GetConnStr(cfg))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer m.Close()

	// if err := m.Up(); err != nil && err != migrate.ErrNoChange {
	// 	fmt.Println(err)
	// 	return
	// }

	// userID, err := uuid.Parse("fcd30c1d-fa2b-4d41-9512-c27c245494ec")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// 	// log.Fatal("Invalid UUID:", err)
	// }
	userID, err := handlers.GetFirstUserId(database)
	if err != nil {
		fmt.Println(err)
		return
	}
	subscribes, err := handlers.GetUserSubscribes(userID, database)
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
