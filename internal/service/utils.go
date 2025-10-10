package service

import (
	"fmt"

	"github.com/google/uuid"
)

// userID, err := uuid.Parse("fcd30c1d-fa2b-4d41-9512-c27c245494ec")
// if err != nil {
// 	fmt.Println(err)
// 	return
// 	// log.Fatal("Invalid UUID:", err)
// }

func ParseUUID(userID string) (uuid.UUID, error) {
	var parsedUserID uuid.UUID

	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		fmt.Println(err)
		return parsedUserID, err
	}

	return parsedUserID, nil
}
