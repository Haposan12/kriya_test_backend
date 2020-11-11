package utils

import guuid "github.com/google/uuid"

func GenerateNewUUID() string {
	uuid := guuid.New()

	return uuid.String()
}