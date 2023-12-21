package main

import (
	"fmt"
	"github.com/lukasgolson/FileArray/fileArray"
	"github.com/lukasgolson/FileArray/serialization"
	"math"
	"math/rand"
	"runtime"
	"sync"
)

func generateRandomNumbers(start, end int, wg *sync.WaitGroup) {
	defer wg.Done()

	fileName := fmt.Sprintf("sequence_%d_%d.bin", start, end)

	fa, err := fileArray.NewFileArray[Number32](fileName, false)
	err = fa.Expand(serialization.Length(end - start))
	if err != nil {
		return
	}

	if err != nil {
		fmt.Println("Error creating file array:", err)
		return
	}

	for i := start; i < end; i++ {
		random := rand.New(rand.NewSource(int64(i)))
		_, err := fa.Append(NewNumber32(random.Intn(100) + 1))
		if err != nil {
			return
		}
	}

	err = fa.Close()
	if err != nil {
		return
	}
}

func main() {
	var numberOfElements = math.MaxInt32

	var wg sync.WaitGroup

	numGoroutines := runtime.NumCPU() - 1

	elementsPerRoutine := numberOfElements / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i * elementsPerRoutine
		end := start + elementsPerRoutine
		if i == numGoroutines-1 {
			end = numberOfElements
		}
		wg.Add(1)
		go generateRandomNumbers(start, end, &wg)
	}

	wg.Wait()
}
