package main

import "fmt"

func DirectionalVsNilChannels() {
	DirectionalChannels()
	NilChannels()
}

/*
a. Directional Channel
By default, channels are bidirectional, permitting data to flow both ways. Go provides an option to restrict them, specifying if a channel should only send or receive data.

But, in technical terms, these aren’t different ‘types’ of channels.

They simply apply a compile-time constraint to ensure that channels operate correctly and are used in the expected way in specific contexts.
*/

func DirectionalChannels() {
	/*
		Send-only Channel
		This channel can only send values, creating a clear communication that certain channels are strictly for sending data:
	*/

	var sendCh chan<- int = make(chan int)
	// v := <-sendCh // invalid operation: cannot receive from send-only channel sendCh

	/*
		Receive-only Channel
		In contrast, this channel can only receive data, blocking any attempts to send data:
	*/
	var recvCh <-chan int = make(chan int)
	// recvCh <- 42 // invalid operation: cannot send to receive-only channel recvCh
}

/*
“If a channel can only send or only receive, how is it useful? Don’t we need both functionalities for it to work?”

A valid point.

The key reason for having directional channels is to enforce particular behaviors in different parts of the code to avoid errors and clarify its purpose, not during channel creation, but during its usage.

When you pass a send-only channel to a function, you’re clearly stating: “This function is designed to produce data and send it, it should not be responsible for receiving data”:
*/

// The function can only send to the channel
func produce(ch chan<- int) {
	ch <- 42
}

// The function can only receive from the channel.
func consumer(ch <-chan int) int {
	return <-ch
}

/*
b. Nil Channel
Channels in Go can hold a nil value, and dealing with them comes with its own set of rules that might feel a bit unusual if you're navigating through Go for the first time.

Send and Receive Actions:
When you try to send data through a nil channel, you might expect to face a panic.

But, instead, what happens is perpetual blocking:
*/

func NilChannels() {
	var ch chan int // ch is nil initially

	go func() {
		ch <- 42 // This will block indefinitely
	}()

	value := <-ch // This too, will block indefinitely
	fmt.Println(value)

	/*
	   There are no panics, no errors. So, you might wonder, is there a useful side to this?

	   Indeed, there are scenarios where this characteristic might be exploited intentionally to deactivate a channel selection, a tactic we’ll uncover more comprehensively in an upcoming discussion.

	   Although sending or receiving through a nil channel won’t raise a panic, attempting to close it will:
	*/

	var ch1 chan int // ch1 is nil by default
	close(ch1)       // panic: close of nil channel
}
