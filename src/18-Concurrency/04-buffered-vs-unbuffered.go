/*
Channel Types
We’re diving into two main channel types in Go: unbuffered and buffered channels.

While some may argue there are more types, particularly when looking at directional channels, I view those as merely syntactical variations, limited by the compiler

Rather than classifying channel types by certain properties, we’ll talk about them through their behavior, including unbuffered, buffered, directional, nil channels, and closed channels.
*/

package main

import "fmt"

func BufferedVsUnbufferedChannels() {
	UnbufferedChannels()
	BufferedChannels()
}

/*
a. Unbuffered Channel
You might recognize this type from earlier examples.

An unbuffered channel is what you get when you initialize a channel without specifying a size, or with a size of 0.
*/

func UnbufferedChannels() {
	uch1 := make(chan int)
	// uch2 := make(chan int, 0) // also an unbuffered channel

	/*
	   Immediate Blocking:
	   When a value is sent on an unbuffered channel (e.g., uch1 <- 42), the send operation is blocked until another goroutine executes a corresponding receive (e.g., value := <-uch1).

	   Synchronization:
	   This blocking behavior makes unbuffered channels naturally synchronize.

	   So, if one goroutine sends data and another receives it, the operations are synchronized at the point of data exchange, known as “synchronous communication”.

	   Deadlock:
	   If a goroutine sends data to an unbuffered channel, there must be another goroutine to receive from the channel, and vice versa, to avoid deadlock.

	   If not, the program will fatal (not panic).
	*/

	fmt.Println("Receiving from channel...")
	val := <-uch1 // fatal error: all goroutines are asleep - deadlock!

	fmt.Println("Received:", val)
}

/*
b. Buffered Channel
Buffered channels introduce an interesting twist… they don’t block immediately upon sending a value. Instead, they can accommodate a specific number of values before they get blocked.

Think of it like a mini storage space within the channel.
*/

func BufferedChannels() {
	ch := make(chan int, 3)

	go func() {
		fmt.Println("Goroutine: Waiting for a value from the channel...")
		fmt.Printf("Goroutine: Got the value %d from the the channel.\n", <-ch)
	}()

	// Populating the buffered channel to its limit.
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println("Buffer is now full.")
	ch <- 4
	fmt.Println("This line prints after the 4th value gets through.")
}

/*
OUTPUT:
Buffer is now full.
Goroutine: Waiting for a value from the channel...
Goroutine: Got the value 1 from the channel.
This line prints after the 4th value gets through.
*/

/*
After pushing 3 values to the channel, we’re able to proceed to the message “Buffer is now full”, but the 4th value makes us pause until the goroutine fetches a value.

“What about when we fetch data? Will it block too?”

Absolutely, the receiving side does block if there’s nothing to fetch.

Capacity & Length:
Buffered channels introduce us to two new concepts: capacity and length.

Capacity: The maximum number of values a channel can store.

Length: The current count of values in the channel

Taking the earlier example, when sending the 4th value with ch <- 4, can you guess len(ch)? Try it yourself before reading on.

If you thought 3, you’re right on the mark, even when full, the channel length stays at 3, pausing the send operation until there’s space.
*/
