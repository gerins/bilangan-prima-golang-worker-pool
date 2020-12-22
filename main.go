package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

var (
	bufferedChannel = 1000
	total           = 0
)

func main() {
	startTime := time.Now()
	totalWorker := 1
	maxPrime := 20000000

	chanDistribution := numberProduction(maxPrime)
	result := dispatchWorker(chanDistribution, totalWorker)

	for range result {
		// fmt.Println(data)
	}

	fmt.Println(`Mencari Prime Number 2 sampai`, maxPrime)
	fmt.Println(`Dengan total Worker =`, totalWorker)
	fmt.Println(`Buffered Channel =`, bufferedChannel)
	fmt.Println(`Total duration =`, time.Since(startTime))
	fmt.Println(`Total data received`, total)
	fmt.Println(`------------------------------------`)
}

func dispatchWorker(job chan int, totalWorker int) chan int {
	chanOut := make(chan int, bufferedChannel)
	wg := new(sync.WaitGroup)
	mtx := new(sync.Mutex)
	wg.Add(totalWorker)

	go func() {
		for i := 0; i < totalWorker; i++ {
			go func(mtx *sync.Mutex, id int) {
				for data := range job {
					mtx.Lock()
					total++
					mtx.Unlock()
					if result := cariPrime(data); result != 0 {
						chanOut <- result
					}
				}
				wg.Done()
			}(mtx, i)
		}
	}()

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}

func numberProduction(maxPrime int) chan int {
	chanOut := make(chan int, bufferedChannel)
	go func() {
		for i := 1; i <= maxPrime; i++ {
			chanOut <- i
		}
		close(chanOut)
	}()
	return chanOut
}

func cariPrime(prime int) int {
	sqrt := int(math.Sqrt(float64(prime)))
	cek := true // buat cek, kalo tetep true berarti prima

	for i := 2; i <= sqrt; i++ { // pengulangan pengecekan
		if prime%i == 0 { // prime di mod i kalo 0 berarti bukan prima
			cek = false // ubah ke false, biar return 0
			break
		}
	}
	if cek == true {
		return prime
	}
	return 0
}
