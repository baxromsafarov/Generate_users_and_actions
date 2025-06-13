package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"Concurrency/generator"
	"Concurrency/saver"
	"Concurrency/user"
)

func main() {
	rand.Seed(time.Now().Unix())

	start := time.Now()
	wg := &sync.WaitGroup{}

	users := make(chan user.User)
	go generator.GenerateUsers(1000, users)

	for u := range users {
		wg.Add(1)
		go saver.SaveUserInfo(u, wg)
	}

	wg.Wait()

	fmt.Println("Time since last update:", time.Since(start).String())
}
