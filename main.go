package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	totalWorker := 40
	maxPrime := 20000000
	job := numberProduction(maxPrime)

	for data := range dispatchWorker(job, totalWorker) {
		fmt.Println(data)
	}

	fmt.Println(`Mencari Prime Number 2 sampai`, maxPrime)
	fmt.Println(`Dengan total Worker =`, totalWorker)
	fmt.Println(`Buffered Channel =`, 1000)
	fmt.Println(`Total duration =`, time.Since(startTime))
}

func dispatchWorker(job chan *int, totalWorker int) chan int {
	chanOut := make(chan int, 1000)
	wg := new(sync.WaitGroup)
	wg.Add(totalWorker)

	go func() {
		for i := 0; i < totalWorker; i++ {
			go func() {
				for data := range job {
					if result := cariPrime(data); result != 0 {
						chanOut <- result
					}
				}
				wg.Done()
			}()
		}
	}()

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}

func numberProduction(maxPrime int) chan *int {
	chanOut := make(chan *int, 1000)

	go func() {
		for i := 2; i < maxPrime; i++ {
			chanOut <- &i
		}
		close(chanOut)
	}()

	return chanOut
}

func cariPrime(prime *int) int {
	sqrt := int(math.Sqrt(float64(*prime)))
	cek := true // buat cek, kalo tetep true berarti prima

	for i := 2; i <= sqrt; i++ { // pengulangan pengecekan
		if *prime%i == 0 { // prime di mod i kalo 0 berarti bukan prima
			cek = false // ubah ke false biar di line 33 ngga ke append
			break       // hentikan loop, balik ke prime ++ ( line 23)
		}
	}
	if cek == true {
		return *prime
	}
	return 0
}
