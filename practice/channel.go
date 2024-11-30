package main

import (
	"fmt"
	"sync"
)

func reader(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		val, ok := <-ch
		if !ok {
			fmt.Printf("channel close")
			return
		}
		fmt.Printf("Reader %d received %d\n", id, val)
	}
}

func main() {

	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(4)
	go reader(1, ch, &wg)
	go reader(2, ch, &wg)
	go reader(3, ch, &wg)
	go reader(4, ch, &wg)

	for i := 0; i < 100; i++ {
		ch <- i
	}

	close(ch)
	wg.Wait()
}
