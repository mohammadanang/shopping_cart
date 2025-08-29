package util

import (
	"fmt"
	"strings"
	"time"
)

func GenerateKeyNumber(s string) string {
	prefix := strings.ToUpper(s[:1]) + s[1:]
	suffix := fmt.Sprintf("%d", time.Now().UnixMilli())

	return prefix + "-" + suffix
}
