// goImpl
package main

import (
	"fmt"
	"sync"
	"time"
)

var i int = 0

func thread_1() {
	for counter := 0; counter < 1000000; counter++ {
		i++

	}
}

func thread_2() {
	for counter := 0; counter < 1000000; counter++ {
		i--
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go thread_1()
	go thread_2()
	time.Sleep(1000 * time.Millisecond)

	fmt.Println(i)
	/* The scheduler suspends and starts the threads at random,
	meaning the first thread could increment to any number < 10^6
	then the second thread starts decrementing it to any
	> i's current value - 10^6 and so on.
	*/

}
