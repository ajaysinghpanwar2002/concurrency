# Concurrency

* Go's Philosophy: Don't communicate by sharing memory, share memory by communicating
* Don't over-engineer things by using shared memory and complicated, error-prone synchronization primitives; instead, use message-passing between GoRoutines so variables and data can be used in the appropriate sequence.
* A golden rule for concurrency : if you don't need it, don't use it.
* Keep your application's complexity to an absolute minimum; it's easier to write, easier to understand, and easier to maintain
