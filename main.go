package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/talgat-ruby/mechta-go/internal/calc"
	"github.com/talgat-ruby/mechta-go/internal/config"
	"github.com/talgat-ruby/mechta-go/internal/load"
)

const BlockSize = 100

func main() {
	conf := config.NewConfig()

	payload, err := load.File(conf.File)
	if err != nil {
		log.Fatalf("error json load %s", err)
	}

	sum := make(chan int)

	go addPayload(payload, conf.Max, sum)

	fmt.Println("total sum:", getTotal(sum))
}

func getTotal(sum <-chan int) int64 {
	var total int64

	for s := range sum {
		total += int64(s)
	}

	return total
}

func addPayload(payload []map[string]int, max int, sum chan<- int) {
	payloadLen := len(payload)
	wg := &sync.WaitGroup{}
	sem := make(chan struct{}, max)

	for i := 0; payloadLen > 0; i++ {
		sem <- struct{}{}
		wg.Add(1)

		var sl []map[string]int

		if payloadLen > BlockSize {
			sl = payload[i*BlockSize : (i+1)*BlockSize]
			payloadLen -= BlockSize
		} else {
			sl = payload[i*BlockSize : i*BlockSize+payloadLen]
			payloadLen = 0
		}

		go addBlock(sl, wg, sem, sum)
	}

	wg.Wait()
	close(sum)
	close(sem)
}

func addBlock(sl []map[string]int, wg *sync.WaitGroup, sem <-chan struct{}, sum chan<- int) {
	defer wg.Done()
	s := calc.SliceMapSum(sl)
	sum <- s
	<-sem
}
