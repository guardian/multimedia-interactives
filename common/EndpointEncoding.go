package common

import (
	"log"

	"github.com/aws/smithy-go/rand"
)

/*
GenerateNumericId generates a unique numeric ID
*/
func GenerateNumericId() int32 {
	randId, err := rand.CryptoRandInt63n(99999)
	if err != nil {
		log.Fatal("Could not generate random number")
	}
	return int32(randId) + 9900000
}
