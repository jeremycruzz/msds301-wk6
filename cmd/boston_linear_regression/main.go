package main

import (
	"fmt"
	"math"
	"time"

	"github.com/jeremycruzz/msds301-wk6/internal/regression"
)

const NUM_RUNS = 100

func main() {
	fmt.Println("Running Synchronous")

	var totalMseA, minMseA, maxMseA float64 = 0, math.Inf(1), 0
	var totalMseB, minMseB, maxMseB float64 = 0, math.Inf(1), 0
	var totalTime, minTime, maxTime int64 = 0, math.MaxInt64, 0

	for i := 0; i < NUM_RUNS; i++ {
		start := time.Now()

		modelA := regression.New(regression.OmitA)
		mseA := modelA.TrainAndPredict()

		modelB := regression.New(regression.OmitB)
		mseB := modelB.TrainAndPredict()

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
