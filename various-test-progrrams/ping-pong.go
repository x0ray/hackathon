// The ping pong example from Sameer Ajmani
// ref: https://www.youtube.com/watch?v=QDDwwePbDtw
package main

import (
	"fmt"
	"time"
)

// OUTPUT
// pong 1
// ping 2
// pong 3
// ping 4
// pong 5
// ping 6
// pong 7
// ping 8
// pong 9
// ping 10
// pong 11
// ping 12
// Program exited.

type Ball struct {
	hits int
}

func main() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)
	table <- new(Ball) // game on; toss the ball
	time.Sleep(time.Second)
	<-table // game over, grab the ball
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(time.Millisecond * 100)
		table <- ball
	}
}
