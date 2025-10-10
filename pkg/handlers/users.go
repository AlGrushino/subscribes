package handlers

import (
	"github.com/AlGrushino/subscribes/pkg/db"
	"github.com/AlGrushino/subscribes/pkg/models"
)

func GetUserByAge(age int) ([]models.Person, error) {
	persons := []models.Person{}

	cfg, err := db.GetConfig()
	if err != nil {
		return nil, err
	}

	db, err := db.DBInit(cfg)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT username, age FROM users WHERE age = $1", age)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var person models.Person
		err := rows.Scan(&person.Name, &person.Age)
		if err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return persons, nil
}
