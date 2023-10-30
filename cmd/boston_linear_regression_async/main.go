package main

import (
	"fmt"
	"math"
	"time"

	"github.com/jeremycruzz/msds301-wk6/internal/regression"
)

const NUM_RUNS = 100

func main() {
	fmt.Println("Running ASynchronous")

	var totalMseA, minMseA, maxMseA float64 = 0, math.Inf(1), 0
	var totalMseB, minMseB, maxMseB float64 = 0, math.Inf(1), 0
	var totalTime, minTime, maxTime int64 = 0, math.MaxInt64, 0

	for i := 0; i < NUM_RUNS; i++ {
		mseChanA := make(chan float64)
		mseChanB := make(chan float64)
		start := time.Now()

		go func() {
			modelA := regression.New(regression.ColumnsA)
			mseChanA <- modelA.TrainAndPredict()
		}()
		go func() {
			modelB := regression.New(regression.ColumnsB)
			mseChanB <- modelB.TrainAndPredict()
		}()

		mseA := <-mseChanA
		mseB := <-mseChanB
		endTime := time.Since(start).Nanoseconds()
		totalTime += endTime

		totalMseA += mseA
		totalMseB += mseB

		// mse A
		if mseA < minMseA {
			minMseA = mseA
		}
		if mseA > maxMseA {
			maxMseA = mseA
		}

		// mse B
		if mseB < minMseB {
			minMseB = mseB
		}
		if mseB > maxMseB {
			maxMseB = mseB
		}

		// time
		if endTime < minTime {
			minTime = endTime
		}
		if endTime > maxTime {
			maxTime = endTime
		}
	}

	avgMseA := totalMseA / NUM_RUNS
	avgMseB := totalMseB / NUM_RUNS
	avgTime := totalTime / NUM_RUNS

	fmt.Printf("Model A ------Min MSE: %v, Max MSE: %v, Avg MSE: %v\n", minMseA, maxMseA, avgMseA)
	fmt.Printf("Model B ------Min MSE: %v, Max MSE: %v, Avg MSE: %v\n", minMseB, maxMseB, avgMseB)
	fmt.Printf("Time    ------Min Time: %v ns, Max Time: %v ns, Avg Time: %v ns\n", minTime, maxTime, avgTime)
}
