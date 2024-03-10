## sync.Mutex

* Mutex = "mutual exclusion" - allows us to deal with race conditions
* Relative simple to use
* Dealing with shared resource and concurrent/parallel goroutines
* Lock/Unlock
* We can test for race conditions when running code, or testing it

## Channels

* Channels are a means of having GoRoutines share data 
* They can talk to each other
* This is Go's philosophy of having things share memory by communicating, rather than communicating by sharing memory.
* The Producer/Consumer problem