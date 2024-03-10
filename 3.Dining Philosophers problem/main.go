package main

import (
	"fmt"
	"sync"
	"time"
)

var philosophers = []string{"Ajay", "Dips", "Vikram", "Dadu", "Vivek"}
var wg sync.WaitGroup
var sleepTime = 1 * time.Second
var eatTime = 2 * time.Second
var thinkTime = 1 * time.Second
var orderFinished []string
var orderMutex sync.Mutex

const hunger = 3

func diningProblem(philosopher string, leftFork, RightFork *sync.Mutex) {
	defer wg.Done()

	fmt.Println(philosopher, "is seated")
	time.Sleep(sleepTime)

	for i := hunger; i > 0; i-- {
		fmt.Println(philosopher, "is hungry.")
		time.Sleep(sleepTime)

		// lock both forks
		leftFork.Lock()
		fmt.Printf("%s picked up the fork to his left.\n", philosopher)

		RightFork.Lock()
		fmt.Printf("%s picked up the fork to his right.\n", philosopher)

		fmt.Println(philosopher, "has both forks, and is eating")
		time.Sleep(eatTime)

		// give philosopher some time to think
		fmt.Println(philosopher, "is thinking")
		time.Sleep(thinkTime)

		// unlock the mutexes
		RightFork.Unlock()
		fmt.Printf("%s put down the fork to his right.\n", philosopher)
		leftFork.Unlock()
		fmt.Printf("%s put down the fork to his left.\n", philosopher)
		time.Sleep(sleepTime)
	}

	// print out done message
	fmt.Println(philosopher, "is satisfied")
	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher)
	orderMutex.Unlock()

	time.Sleep(sleepTime)

	fmt.Println(philosopher, "has left the table.")
}

func main() {
	fmt.Println("The dining philosopher problem")
	fmt.Println("------------------------------")

	var numberOfPhilosophers int = len(philosophers)
	wg.Add(numberOfPhilosophers)
	forkLeft := &sync.Mutex{}

	for i := 0; i < numberOfPhilosophers; i++ {
		// create a mutes for rigth fork
		forkRight := &sync.Mutex{}

		go diningProblem(philosophers[i], forkLeft, forkRight)
		forkLeft = forkRight
	}

	wg.Wait()
	fmt.Println("The table is empty.")
	fmt.Println("-------------------")

	for i := 0; i < numberOfPhilosophers; i++ {
		fmt.Printf("%s got the %d position in eating food \n", orderFinished[i], i+1)
	}
}
