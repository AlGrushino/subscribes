package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/AlGrushino/subscribes/internal/repository/models"
	"github.com/google/uuid"
)

func handleDates(subscription *models.Subscribe, startDate, endDate string) error {
	if startDate == "" {
		return errors.New("parameter startDate is required")
	}

	parsedStartDate, err := time.Parse("01-2006", startDate)
	if err != nil {
		return fmt.Errorf("failed to parse startDate, original error: %w", err)
	}
	subscription.StartDate = parsedStartDate

	if endDate == "" {
		// мб здесь будут проблемы
		subscription.EndDate = nil
		return nil
	}

	parsedEndDate, err := time.Parse("01-2006", endDate)
	if err != nil {
		return fmt.Errorf("failed to parse endDate, original error: %w", err)
	}
	subscription.EndDate = &parsedEndDate

	return nil
}

func handlePrice(subscription *models.Subscribe, price int) error {
	if price < 0 {
		return errors.New("price must not be negative")
	}
	subscription.Price = price
	return nil
}

func handleServiceName(subscription *models.Subscribe, name string) error {
	if name == "" {
		return errors.New("serviceName must not be empty")
	}
	subscription.ServiceName = name
	return nil
}

func handleUserID(subscription *models.Subscribe, userID string) error {
	if userID == "" {
		return errors.New("userID must not be empty")
	}

	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	subscription.UserUUID = parsedUserID
	return nil
}
