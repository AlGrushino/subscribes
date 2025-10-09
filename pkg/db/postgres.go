package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

func MakeQuery() ([]Person, error) {
	res := []Person{}
	connStr := "postgres://postgres:postgres@localhost:5432/subscribe_project"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, errors.New("бд не открылась")
	}
	defer db.Close()

	age := 21
	rows, err := db.Query("SELECT username, age FROM users WHERE age = $1", age)

	if err != nil {
		return nil, fmt.Errorf("ошибка запроса: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var person Person
		err := rows.Scan(&person.Name, &person.Age)
		if err != nil {
			return nil, errors.New("не удалось отсканировать персону")
		}
		res = append(res, person)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.New("произошла какая-то дич")
	}
	return res, nil
}

type Person struct {
	Name string
	Age  int
}
