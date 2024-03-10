package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	msg = s
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go updateMessage("Hello, Universe", &wg)
	go updateMessage("Hello, Ajay", &wg)
	wg.Wait()

	fmt.Println(msg)
}

/*
Using mutex, It is handling the race condition. the above one is not handling. it just the basic waitgroup.

package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	msg = s
	m.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hello, Universe", &wg, &mutex)
	go updateMessage("Hello, Ajay", &wg, &mutex)
	wg.Wait()

	fmt.Println(msg)
}

*/
