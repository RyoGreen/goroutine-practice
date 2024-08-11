package main

import "fmt"

func goroutine(s []string, c chan string) {
	var keyword string
	for _, v := range s {
		keyword += v
		c <- keyword
	}
	close(c)
}

func main() {
	words := []string{"test1!", "test2!", "test3!", "test4!"}
	c := make(chan string)
	go goroutine(words, c)
	for w := range c {
		fmt.Println(w)
	}
}
