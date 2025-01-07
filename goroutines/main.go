package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	response := make(chan any, 3)
	wg := &sync.WaitGroup{}

	wg.Add(3)

	go fetchUser(response, wg)
	go fetchUserLike(response, wg)
	go fetchUserMatch(response, wg)

	go func() {
		wg.Wait()
		close(response)
	}()

	for res := range response {
		fmt.Println("res :", res)
	}

	fmt.Println("Took :", time.Since(start))
}

func fetchUser(response chan any, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Millisecond * 100)

	response <- "Rudrik"
}

func fetchUserLike(response chan any, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Millisecond * 150)

	response <- 1000
}

func fetchUserMatch(response chan any, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Millisecond * 200)

	response <- "Prajapati"
}
