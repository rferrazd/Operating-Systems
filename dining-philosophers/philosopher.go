package main

import (
	"fmt"
	"html"
	"strconv"
	"sync"
)

const (
	FOOD   = "0x0001F35C"
	FINISH = "0x0001F44C"
)

// MAKE A PHOLOSOPHER STRUCT. give it an ID and name
type Philosopher struct {
	ID             int // ID = what process number it is
	Name           string
	LeftChopStick  *ChopStick // pointer so actually a memory access
	RightChopStick *ChopStick
}

// this function operation on a philosopher type
// func (type you can call function on) Name() return type
func (p *Philosopher) Eat(wg *sync.WaitGroup) { // pointer is on the type
	// functions available with Mutex: lock and unlock
	wg.Add(1) // one go routine has started
	// philosopher "one group of threads"
	p.LeftChopStick.Lock() // .Lock() part of sync mutex package
	p.RightChopStick.Lock()

	fmt.Printf("%d is eating %s\n", p.ID, GetEmoticon(FOOD))

	// Unlock chopsticks
	p.LeftChopStick.Unlock() // public function starts with capital
	p.RightChopStick.Unlock()

	// perfect for deadlock state
	fmt.Printf("%d is done eating %s\n", p.ID, GetEmoticon(FINISH))
	// add more for resolving deadlock state

	wg.Done() // threat finsihed decrement, get to zero
}

// name of parameter then type
func GetEmoticon(value string) string {
	i, _ := strconv.ParseInt(value, 0, 64) // use strconv to parse an int from a string
	foodEmoticon := html.UnescapeString(string(i))
	return foodEmoticon
}
