package utils

import (
	"strings"

	"github.com/google/uuid"
)

func NewHostID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
