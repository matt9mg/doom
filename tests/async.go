package main

import (
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	semaphore := make(chan struct{}, 0)

	for {
			wg.Add(1)
			defer wg.Done()

			go func() {
				defer wg.Done()
				wg.Add(1)
				semaphore <- struct{}{} // Lock
				func() {
					log.Println("here")
					<-semaphore // Unlock
				}()
		}()
	}

	wg.Wait()
}