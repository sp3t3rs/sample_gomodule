package uuid_generator

import (
	"github.com/google/uuid"
)

func GenerateUUID() UUID {
	return uuid.New()
}
