package common

import (
	"fmt"
	"time"
)

func Now() string {
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()
	return fmt.Sprintf("%d:%d:%d", hour, minute, second)
}

func LogNow() {
	fmt.Println(Now())
}