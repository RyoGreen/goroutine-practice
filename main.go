package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	userName := fetchUser()
	respLikeCh := make(chan int, 1)
	respMatchCh := make(chan string, 1)
	var wg sync.WaitGroup
	wg.Add(2)
	go fetchUserLikes(userName, respLikeCh, &wg)
	go fetchUserMatch(userName, respMatchCh, &wg)
	wg.Wait()

	func() {
		close(respLikeCh)
		close(respMatchCh)
	}()

	fmt.Println("likes: ", <-respLikeCh)
	fmt.Println("match: ", <-respMatchCh)
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 500)
	return "Taro"
}

func fetchUserLikes(userName string, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 450)
	ch <- 1
}

func fetchUserMatch(userName string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 200)
	ch <- "Hanako"
}
