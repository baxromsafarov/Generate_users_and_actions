package user

import "fmt"

type User struct {
	ID    int
	Email string
	Logs  []LogItem
}

func (u User) GetActivityInfo() string {
	out := fmt.Sprintf("ID: %d | Email: %s\nActivity Log:\n", u.ID, u.Email)
	for i, item := range u.Logs {
		out += fmt.Sprintf("%d. [%s] at %s\n", i+1, item.Action, item.Timestamp)
	}
	return out
}
