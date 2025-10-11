package service

import (
	"fmt"

	"github.com/google/uuid"
)

func ParseUUID(userID string) (uuid.UUID, error) {
	var parsedUserID uuid.UUID

	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		fmt.Println(err)
		return parsedUserID, err
	}

	return parsedUserID, nil
}
