package main

import "sync"

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	m := make(map[int]int)

	for i := 0;i < 10;i++ {
		go func() {
			m[i] = i
			wg.Done()
		}()
	}
	wg.Wait()

}

