/*
In concurrent programming, a Mutex (Mutual Exclusion) is a locking mechanism that prevents race conditions by ensuring only one goroutine can access a shared resource at a time. It's akin to a key to a room – only one person can hold the key and enter at once.
*/
package main

import (
	"fmt"
	"sync"
)

/*
EXAMPLE 1: The Problem Without a Mutex
Multiple goroutines increment a counter at the same time.
Without a mutex, they step on each other → wrong result.
*/

func withoutMutex() {
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // RACE CONDITION: multiple goroutines read+write at the same time
		}()
	}

	wg.Wait()
	fmt.Println("[No Mutex] Expected 1000, got:", counter) // likely NOT 1000
}

/*
EXAMPLE 2: The Fix — Basic Mutex
Lock() before touching shared data, Unlock() when done.
Only one goroutine can be between Lock/Unlock at a time.
*/

func withMutex() {
	counter := 0
	var mu sync.Mutex // the "key to the room"
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()   // grab the key — other goroutines will wait here
			counter++   // safe: only one goroutine is here at a time
			mu.Unlock() // give the key back
		}()
	}

	wg.Wait()
	fmt.Println("[With Mutex] Expected 1000, got:", counter) // always 1000
}

/*
EXAMPLE 3: Using defer for Unlock (Best Practice)
defer mu.Unlock() runs when the function returns,
so you can never accidentally forget to unlock.
*/

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock() // always unlocks, even if the function panics
	c.value++
}

func (c *SafeCounter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func withDeferUnlock() {
	counter := &SafeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("[Defer Unlock] Counter:", counter.Get()) // always 1000
}

/*
EXAMPLE 4: RWMutex — Faster Reads, Safe Writes
Use RWMutex when you have many readers but few writers.
Multiple goroutines can read at the same time (RLock),
but a write (Lock) gets exclusive access.
*/

type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

func (m *SafeMap) Set(key string, value int) {
	m.mu.Lock() // exclusive lock for writing
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *SafeMap) Get(key string) int {
	m.mu.RLock() // shared lock — multiple readers allowed simultaneously
	defer m.mu.RUnlock()
	return m.data[key]
}

func withRWMutex() {
	sm := &SafeMap{data: make(map[string]int)}
	var wg sync.WaitGroup

	// One writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		sm.Set("score", 42)
	}()

	// Many concurrent readers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("[RWMutex] score =", sm.Get("score"))
		}()
	}

	wg.Wait()
}

/*
EXAMPLE 5: TryLock — Non-Blocking Attempt (Go 1.18+)
TryLock returns false immediately if the mutex is already
locked, instead of blocking. Useful for "do it if you can,
skip it if you can't" patterns.
*/

func withTryLock() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	mu.Lock() // lock held by "main"

	for i := 0; i < 3; i++ {
		wg.Add(1)
		id := i
		go func() {
			defer wg.Done()
			if mu.TryLock() {
				fmt.Printf("[TryLock] goroutine %d acquired the lock\n", id)
				mu.Unlock()
			} else {
				fmt.Printf("[TryLock] goroutine %d could not acquire the lock, skipping\n", id)
			}
		}()
	}

	wg.Wait()
	mu.Unlock() // release the original lock
	fmt.Println("[TryLock] done")
}

func Mutexes() {
	fmt.Println("=== Example 1: Without Mutex (race condition) ===")
	withoutMutex()

	fmt.Println("\n=== Example 2: Basic Mutex ===")
	withMutex()

	fmt.Println("\n=== Example 3: Defer Unlock (best practice) ===")
	withDeferUnlock()

	fmt.Println("\n=== Example 4: RWMutex (read-heavy workloads) ===")
	withRWMutex()

	fmt.Println("\n=== Example 5: TryLock (non-blocking) ===")
	withTryLock()
}
