package common

import "github.com/google/uuid"

func UuidToString(uuid uuid.UUID) string {
	return uuid.String()
}
