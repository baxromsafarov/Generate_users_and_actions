package saver

import (
	"Concurrency/user"
	"fmt"
	"os"
	"sync"
	"time"
)

func SaveUserInfo(user user.User, wg *sync.WaitGroup) error {
	defer wg.Done()

	time.Sleep(time.Millisecond * 10)
	fmt.Printf("WRITING FILE FOR USER ID : %d\n", user.ID)

	filename := fmt.Sprintf("logs/user-%d.txt", user.ID)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(user.GetActivityInfo())
	if err != nil {
		return err
	}

	return nil
}
