package tools

import (
	"fmt"
	"time"
)

func MysqlDate() string{
	t:= time.now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
	t.Year(), t.Month(),t.Day(), t.Hour(),t.Minute(),t.Second()
	)
}