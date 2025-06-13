package generator

import (
	"Concurrency/user"
	"fmt"
	"math/rand"
	"time"
)

func GenerateUsers(count int, users chan user.User) {
	for i := 0; i < count; i++ {
		users <- user.User{
			ID:    i + 1,
			Email: fmt.Sprintf("user%d@ninja.go", i+1),
			Logs:  GenerateLogs(rand.Intn(1000)),
		}
		//time.Sleep(time.Millisecond * 10)
	}
	close(users)
}

func GenerateLogs(count int) []user.LogItem {
	logs := make([]user.LogItem, count)

	for i := 0; i < count; i++ {
		logs[i] = user.LogItem{
			Timestamp: time.Now(),
			Action:    user.Actions[rand.Intn(len(user.Actions)-1)],
		}
	}
	return logs
}
