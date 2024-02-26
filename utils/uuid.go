package utils

import (
	"github.com/google/uuid"
)

const UUIDv5Namespace = "cf967573-0c1f-43be-8d61-8fecf122a28c"

// GenerateUUIDv5 generates a UUID version 5 based on the provided name.
func GenerateUUIDv5(name string) (uuid.UUID, error) {
	namespace, err := uuid.Parse(UUIDv5Namespace)
	if err != nil {
		return uuid.Nil, err // Return an error if the namespace UUID is invalid
	}
	return uuid.NewSHA1(namespace, []byte(name)), nil
}
