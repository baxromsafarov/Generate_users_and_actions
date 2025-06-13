package user

import "time"

var Actions = []string{
	"logged in",
	"logged out",
	"create record",
	"update record",
	"delete record",
}

type LogItem struct {
	Action    string
	Timestamp time.Time
}
