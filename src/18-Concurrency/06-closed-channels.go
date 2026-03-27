/*
Closed Channel
Closed channels in Go can be tricky, yet understanding their behaviors is quite crucial to effectively manage potential issues.
*/

package main

import "fmt"

func ClosedChannels() {
	/*
		Send Behavior First, let’s talk about sending to a closed channel. Unlike a nil channel, it doesn’t get blocked forever.

		Instead, it results in an immediate panic to prevent unintended sends after the channel has signaled no more data will be sent:
	*/

	ch := make(chan int)
	close(ch)
	ch <- 42 // "send on closed channel"

	/*
		Received Behavior
		On the other hand, receiving from a closed channel is slightly more nuanced:

		If the channel buffer still holds values, they can be received as if the channel is still open.

		Once the buffer is empty, further receives will return the zero value for the channel’s type without blocking.
	*/

	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch)
	fmt.Println(<-ch1) // 1
	fmt.Println(<-ch1) // 2
	fmt.Println(<-ch1) // 3

	// If you’d like to know if a channel has been closed, you can utilize a two-value receive. If the second value (a boolean) is false, the channel is [closed] and [has no values left to receive]:

	ch3 := make(chan int)
	close(ch3)
	value, ok := <-ch3
	fmt.Println(value, ok) // 0 false

	// Similar to sending to a closed channel, trying to close an already closed channel also results in a panic.
	close(ch3) // panic: close of closed channel
}
