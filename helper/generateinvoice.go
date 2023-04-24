package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateInvoice() string {
	now := time.Now().Format("20060102")
	return fmt.Sprintf("MT%s%d", now, rand.Intn(100))
}
