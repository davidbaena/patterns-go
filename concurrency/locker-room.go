package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Player struct {
	id int
}

func (p Player) enterLockerRoom(lockerRoom chan int) {
	lockerRoom <- p.id
	fmt.Printf("Player %d is entering the locker room.\n", p.id)
}

func (p Player) leaveLockerRoom(lockerRoom chan int, n int) {
	fmt.Printf("Player %d has left the locker room after %d seconds.\n", p.id, n)
	<-lockerRoom
}

func main() {
	var wg sync.WaitGroup
	lockerRoom := make(chan int, 3)  // The locker room can hold only 1 player at a time.
	rand.Seed(time.Now().UnixNano()) // Initialize the random number generator.

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			p := Player{id: id}
			p.enterLockerRoom(lockerRoom)
			n := rand.Intn(3) + 1                      // Generate a random integer between 1 and 3.
			time.Sleep(time.Duration(n) * time.Second) // Sleep for the random duration.

			p.leaveLockerRoom(lockerRoom, n)
		}(i)
	}

	wg.Wait()
	fmt.Printf("All players have left the locker room.")
}
