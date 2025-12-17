package helper

import (
	"errors"

	"github.com/google/uuid"
)

func StringToUUID(id string) (uuid.UUID, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, errors.New("invalid uuid format")
	}

	return parsedID, nil
}
