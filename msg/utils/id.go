package utils

import (
	"strings"

	"github.com/google/uuid"
)

func NewEventID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
