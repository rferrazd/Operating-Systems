package main

// HOMEWORK: GO INTO THE MANAGE FUCNTION
// TODO: PPRITN GOLBAL VIEW OF THE TABLE
// go build
// go run
import (
	"fmt"
	"sync"
)

// define constants
const NUM_PHILOSOPHERS int = 5
const NUM_CHOPSTICKS int = 5
const EAT_TIMES int = 3
const NUM_EATING_PHILOSOPHER int = 2

// banker method you have an object that keeps track of the current state of the world. Is that the host?
// Host is banker, keep track of who is eating, requestion
type Host struct {
	// channel for allowing a philosopher to eat
	requestChannel chan *Philosopher // type channel which is a channel of philosopher

	// channel that quits the program; tells the host to stop hosting
	quitChannel chan int

	// keep track/bookkeeping of currently eating philosophers
	eatingPhilosophers map[int]bool // True or False, eating or not

	// we need to lock the bookkeeping variable
	mu sync.Mutex
}

//func (h *Host) manage() {
// doing an infinite loop
//	for {
//			if len(h.requestChannel) == NUM_EATING_PHILOSOPHER {
//				finished := <-h.requestChannel // Pops a Philosopher object off the channel .// h.requestChannel= channel
//				currentlyEating := make([]int, 0 , NUM_PHILOSOPHERS)
//				// eatingPhilosophers is a map has key value pair
//				for index, eating := range h.eatingPhilosophers {
//					if eating {
//						currentlyEating = append(currentlyEating, index) // if they are eating append index to the array.
//					}
//
//				}
//				fmt.Printf("%v have been eating, clearing plates %d\n", currentlyEating, finished.ID) // %d is decimal integer
// finished.ID looking at ID of the philosophers
//				h.eatingPhilosophers[finished.ID] = false  // this particular philosopher is not eating
//			}

// similar to a switch stmt
//		select {
//			case <-h.quitChannel:  // arrow key means that we are watching the channel. this is syntax for blocking on a producing any signal
//					// when the channel received a signal
// end the host manage function
//					fmt.Println("party is over")
//					return
//		default:
//		}
//	}
//}

func main() {
	var wg sync.WaitGroup
	//requestChannel := make( chan *Philosopher, NUM_EATING_PHILOSOPHER)
	//quitChannel := make(chan int, 1)
	//host := Host {
	//	requestChannel: requestChannel,
	//	quitChannel: quitChannel,
	//	eatingPhilosophers: make(map[int]bool)
	//}

	// chopstick array
	chopsticks := make([]*ChopStick, NUM_CHOPSTICKS)
	for c := 0; c < NUM_CHOPSTICKS; c++ {
		chopsticks[c] = &ChopStick{
			ID: c + 1, // syntax for attributes
		}
	}

	// make array of philosopher objects
	philos := make([]*Philosopher, NUM_PHILOSOPHERS)
	// make chopstick array

	for i := 0; i < NUM_PHILOSOPHERS; i++ {
		philos[i] = &Philosopher{
			ID:             i + 1,
			LeftChopStick:  chopsticks[i],
			RightChopStick: chopsticks[(i+1)%5], // wrap around lists circular table
		}
	}
	// create a thread
	//go host.manage() // pulls for channels
	// weight groups used to count the number of philosophers then you signal (wg.Done
	// )
	// create a new go routine (thread) and in the thread run the philosopher eat function
	for _, philosopher := range philos {
		go philosopher.Eat(&wg) // get memory address of the variable
		fmt.Println("Started go routine: ", philosopher.ID)
	}
	fmt.Println("Waiting for the philosopher")
	wg.Wait()
	fmt.Println("Everyone finished eating")
	//<-host.quitChannel
	//}
}

// THIS IS DEADLOCKING CODE. PRINT
// GO IS TO CHANGE SOME STUFF TO FIX DEADLOCK
// diff ~/path to dinner-party.go dinner-party.go
