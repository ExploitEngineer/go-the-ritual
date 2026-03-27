/*
Imagine a channel as a simple pipe, this pipe connects different goroutines, allowing them to talk to each other by sending data into one end of the pipe and receiving it from the other.

Why exactly do we use channels?

Channels give us a safe way for goroutines to communicate and also to keep their actions in sync.

While the usual method of using locks to control access to data can be tricky and might cause deadlocks, channels make managing concurrency straightforward.

They follow a simple idea: ‘Don’t communicate by sharing memory; share memory by communicating.’
*/

package main

import "fmt"

func Channels() {
	// Creating Channels
	// var c1 chan int - another way of creating the unbuffered channel
	ch := make(chan int)

	// Here, c1 is declared as a channel that deals with integers, yet it’s uninitialized, it holds a nil value. In contrast, ch is an unbuffered channel for integers.

	// Channel Basic Operations
	ch <- 42      // sends 42 to channel ch
	value := <-ch // receive from channel ch
	fmt.Println(value)
	close(ch)

	// When you write ch <- 42, the goroutine takes a pause, waiting until a receiver is ready. And it works the same when you’re on the receiving end of a channel.

	closeChannel()
}

// Here’s a simple example to show what’s happening:
func channelExample() {
	ch := make(chan int) // creating an unbuffered channel

	// A goroutine to send a value.
	go func() {
		fmt.Println("Ready to send 42...")
		ch <- 42
		fmt.Println("42 is sent.")
	}()

	// Waiting to get the value from the channel.
	fmt.Println("Waiting for value from channel...")
	val := <-ch
	fmt.Println("Value received:", val)
}

/*
“But what happens when we close a channel?”

Invoking close(ch) gives a clear signal: this channel is no longer in use, if you try to send a value through it, the application will panic.

But interestingly, trying to receive from a closed channel won’t trigger a panic. Instead, it provides the zero value of the channel’s type (assuming it’s unbuffered).
*/

/*
A few things to note:

Closing a channel is a “definitive” act, there’s no turning back or reopening it.

Sadly, Go doesn’t offer a built-in method to check if a channel is closed. You’ll only discover it’s closed when you try to read from it and you get a zero value when it’s void of data.
*/
func closeChannel() {
	ch := make(chan int)
	close(ch)

	v, ok := <-ch

	fmt.Println(v, ok) // 0 false
}

// If the second value (a boolean) is false,
// the channel is [closed] and [has no value left to receive]
