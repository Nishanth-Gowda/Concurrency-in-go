package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	elapsed_time := time.Now()
	userId := 12

	//Using the buffered channel
	response_Ch := make(chan string, 128)

	//Provides utilities that waits till all the goroutines are finished executing
	wg := &sync.WaitGroup{}

	go fetchUserData(userId, response_Ch, wg)
	go fetchUserRecommendations(userId, response_Ch, wg)
	go fetchUserLikes(userId, response_Ch, wg)
	wg.Add(3)
	wg.Wait()

	// Telling Go that this channel will not be used anymore.
	close(response_Ch)

	//fmt.Println(userData)
	// fmt.Println(userRecs)
	// fmt.Println(userLikes)

	for resp := range response_Ch {
		fmt.Println(resp)
	}
	fmt.Println(time.Since(elapsed_time))
}

func fetchUserData(userId int, respCh chan string, wg *sync.WaitGroup) {
	time.Sleep(80 * time.Millisecond)
	respCh <- "user_data........"

	wg.Done()
}

func fetchUserRecommendations(userId int, respCh chan string, wg *sync.WaitGroup) {
	time.Sleep(120 * time.Millisecond)
	respCh <- "user_recommendations"

	wg.Done()
}

func fetchUserLikes(userId int, respCh chan string, wg *sync.WaitGroup) {
	time.Sleep(50 * time.Millisecond)
	respCh <- "user Likes /-/-/-"

	wg.Done()
}