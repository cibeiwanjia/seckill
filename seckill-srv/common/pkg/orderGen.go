package pkg

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func OrderGen(pre string) string {
	timeStr := time.Now().Unix()
	uuidStr := uuid.NewString()[:8]
	orderSn := fmt.Sprintf("%s:%s:%s", pre, timeStr, uuidStr)
	return orderSn
}
