/*
The sync package in Go helps you manage multiple parts of your program that might run at the same time (concurrently) and need to coordinate with each other to avoid conflicts or errors. It provides tools like locks (mutexes), conditions, and wait groups to make sure that different parts of your program don't interfere with each other when they're accessing shared resources or performing tasks simultaneously. This helps ensure that your program runs smoothly and safely even in a multi-threaded or concurrent environment.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func SyncPackage() {}

/*
sync.Mutex

A sync.Mutex in Go is like a lock that one part of your program can hold while it's doing something important, to prevent other parts from interfering. It ensures that only one thing can happen at a time.

Real-life Example: Bank Account Transactions
*/

var (
	balance int
	mutex   sync.Mutex
)

func deposit(amount int) {
	mutex.Lock()
	defer mutex.Unlock()
	balance += amount
	fmt.Printf("Deposited $%d. New balance: %d\n", amount, balance)
}

func withdraw(amount int) {
	mutex.Lock()
	defer mutex.Unlock()
	if balance >= amount {
		balance -= amount
		fmt.Printf("Withdrawn $%d. New balance: $%d\n", amount, balance)
	} else {
		fmt.Println("Insufficient funds.")
	}
}

func syncMutex() {
	balance = 1000

	go deposit(200)
	go withdraw(300)
	go deposit(500)

	time.Sleep(time.Second) // Allow goroutines to finish before exiting
}

/*
In this example, multiple goroutines are simulating deposits and withdrawals from a bank account. The sync.Mutex ensures that only one operation (either deposit or withdrawal) can modify the balance at a time, preventing conflicts.
*/

/*
sync.RWMutex

A sync.RWMutex in Go is like a special lock that allows multiple parts of your program to read data at the same time but ensures that only one part can write (modify) data at a time. It's useful when you have data that is read frequently but modified less often.

Real-life Example: Book Inventory
*/

type Library struct {
	books          map[string]int
	inventoryMutex sync.RWMutex
}

func (l *Library) checkoutBook(title string) {
	l.inventoryMutex.RLock()
	defer l.inventoryMutex.RUnlock()

	// Simulate reading (checking out) a book
	fmt.Printf("Reader checking out book '%s'\n", title)

	remainingCopies := l.books[title]
	if remainingCopies > 0 {
		l.books[title]--
		fmt.Printf("Book '%s' checked out successfully. Remaining copies: %d\n", title)
	} else {
		fmt.Printf("Book '%s' is out of stock.\n", title)
	}
}

func (l *Library) updateInventory(title string, copies int) {
	l.inventoryMutex.Lock()
	defer l.inventoryMutex.Unlock()

	// Simulate updating (restocking) the book inventory
	fmt.Printf("Librarian updating inventory for book '%s'\n", title)
	l.books[title] += copies
	fmt.Printf("Inventory updated for book '%s'. New copies: %d\n", title, l.books[title])
}

func syncRWMutex() {
	library := &Library{
		books: map[string]int{
			"Introduction to Go":    5,
			"Concurrency in Go":     3,
			"Data Structures in Go": 7,
		},
	}

	// Readers (concurrent checkouts)
	for i := range 5 {
		go func(readerID int) {
			time.Sleep(time.Millisecond * time.Duration(readerID)) // Simulate staggered checkouts
			library.checkoutBook("Introduction to Go")
		}(i)
	}

	time.Sleep(time.Second) // Allow some time for readers to start before updating inventory

	// Write (updating inventory)
	library.updateInventory("Introduction to Go", 2)

	time.Sleep(time.Second) // Allow some time for readers and writers to finish before exiting
}

/*
In this example, the sync.RWMutex is used to coordinate access to a shared resource, such as a book inventory in a library. It allows multiple goroutines to read the inventory concurrently using RLock, ensuring efficient parallel reads. For exclusive updates, Lock is employed, guaranteeing that only one goroutine can modify the inventory at a time, preventing conflicts. This ensures a balance between efficient concurrent access for reading and exclusive access for writing, optimizing the program’s
*/

/*
sync.Once

sync.Once in Go is like a guard that ensures a specific action is performed only once, no matter how many times it's called. It's useful for one-time initializations or tasks that should happen exactly once in your program.

Real-life Example: Server Configurations
*/

// Configuration represents the server configuration
type Configuration struct {
	Port    int
	IsDebug bool
	// ...other configuration parameters
}

var (
	serverConfig Configuration
	once         sync.Once
)

// initializeServerConfig simulates the one-time initialization of server configuration.
func initializeServerConfig() {
	fmt.Println("Initializing server configuration...")
	serverConfig = Configuration{
		Port:    8080,
		IsDebug: false,
		// ... initialize other configuration parameters
	}
}

// getServerConfig returns the server configuration, initializing it if not done already.
func getServerConfig() Configuration {
	once.Do(initializeServerConfig)
	return serverConfig
}

func SyncOnce() {
	// In different parts of your program, you might need the server configuration.
	// The configuration will be initialized only once.

	config1 := getServerConfig()
	config2 := getServerConfig()

	fmt.Println("Config1:", config1)
	fmt.Println("Config1:", config2)
}

/*
In this example, initializes a server configuration using sync.Once, ensuring the configuration is set up only once. The getServerConfig function retrieves the configuration, initializing it if needed.
*/
