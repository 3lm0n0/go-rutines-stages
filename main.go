package main

import (
	"fmt"
	"sync"
	"time"
)

func getUsername() string {
	time.Sleep(time.Millisecond * 100)
	
	return "Messi"
}	

func matchUserName(username string, rc chan int, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	
	if username == "Messi" {
		rc <- 1
	}

	close(rc)
	wg.Done()
}

func getUserLikes(rc chan int, rcLikes chan int, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	for l := range rc {
		if l == 1 {
			rcLikes <- 31231
		}
	}
	
	close(rcLikes)
	wg.Done()
}

func print(rcLikes chan int, wg *sync.WaitGroup) {
	for n := range rcLikes {
		if n > 0 {
			fmt.Printf("-> Has %v likes.\n", n)
		} else {
			fmt.Println("-> Not registered user name.")
		}
	}

	wg.Done()
} 

func main() {
	start := time.Now()

	username := getUsername()
	fmt.Printf("-> User name is %s.\n-> Took %v to get.\n", username, time.Since(start))

	wg := sync.WaitGroup{}
	wg.Add(3)
	rcMatch := make(chan int, 1) // response channel fo matching user names.
	rcLikes := make(chan int, 1) // response channel for likes.
	
	go matchUserName(username, rcMatch, &wg) // stage 1
	go getUserLikes(rcMatch, rcLikes, &wg) // stage 2
	go print(rcLikes, &wg)

	wg.Wait()

	fmt.Printf("-> Total time taken %v\n", time.Since(start))
}
