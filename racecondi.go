package main

import (
	"fmt"
	"sync"
)

func race() {
	var wg sync.WaitGroup

	a := 0
	times := 10000

	//萬一沒做lock，就會發生 race condition
	var m sync.Mutex

	wg.Add(times)
	for i := 0; i < times; i++ {
		go func() {
			m.Lock()
			a++
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Printf("a = %d \n", a)
}
/*
鎖的缺點
	效能
	上面的例子用 mutex 來防止多個 goroutine 同時存取同一個變數，因為總共有一萬個 goroutine，
	當你有其中一個正在存取 a 時其他 9999 個都在等他，他們之間完全沒有並行（parallelism），不如用個迴圈把它從 0 加到 10000 可能還更快
	所以在使用鎖時一定要非常小心，只在必要的時候使用，否則效能將會大打折扣
*/
