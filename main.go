package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Level struct {
	x, y          int
	height, level float64
}

type PriorityQueue []Level

func (h PriorityQueue) Len() int           { return len(h) }
func (h PriorityQueue) Less(i, j int) bool { return h[i].level < h[j].level }
func (h PriorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PriorityQueue) Push(x interface{}) {
	*h = append(*h, x.(Level))
}

func (h *PriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func initializeHeap(input [][]float64) *PriorityQueue {
	priorityQueue := &PriorityQueue{}
	heap.Init(priorityQueue)
	for i, v := range input[0] {
		heap.Push(priorityQueue, Level{i, 0, v, v})
	}

	if len(input) > 1 {
		for i, v := range input[len(input)-1] {
			heap.Push(priorityQueue, Level{i, len(input) - 1, v, v})
		}

		if len(input) > 2 {
			for i := 1; i < len(input)-1; i++ {
				heap.Push(priorityQueue, Level{0, i, input[i][0], input[i][0]})
				heap.Push(priorityQueue, Level{len(input[0]) - 1, i, input[i][len(input[0])-1], input[i][len(input[0])-1]})
			}
		}
	}

	return priorityQueue
}

func fillPool(input [][]float64) int {
	height, width := len(input), len(input[0])

	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
		for q := range visited[0] {
			if i == 0 || i == len(visited)-1 || q == 0 || q == len(visited[0])-1 {
				visited[i][q] = true
			}
		}
	}

	result := make([][]float64, height)
	for i := range result {
		result[i] = make([]float64, width)
		for q := range result[i] {
			if i == 0 || i == len(visited)-1 || q == 0 || q == len(visited[0])-1 {
				result[i][q] = input[i][q]
			} else {
				result[i][q] = math.MaxFloat64
			}
		}
	}

	queue := initializeHeap(input)

	for queue.Len() > 0 {
		v := heap.Pop(queue).(Level)
		visited[v.y][v.x] = true
		if v.y > 0 {
			if !visited[v.y-1][v.x] {
				result[v.y-1][v.x] = math.Max(input[v.y-1][v.x], math.Min(result[v.y-1][v.x], v.level))
				heap.Push(queue, Level{v.x, v.y - 1, input[v.y-1][v.x], result[v.y-1][v.x]})
			}
		}
		if v.y < len(input)-1 {
			if !visited[v.y+1][v.x] {
				result[v.y+1][v.x] = math.Max(input[v.y+1][v.x], math.Min(result[v.y+1][v.x], v.level))
				heap.Push(queue, Level{v.x, v.y + 1, input[v.y+1][v.x], result[v.y+1][v.x]})
			}
		}
		if v.x > 0 {
			if !visited[v.y][v.x-1] {
				result[v.y][v.x-1] = math.Max(input[v.y][v.x-1], math.Min(result[v.y][v.x-1], v.level))
				heap.Push(queue, Level{v.x - 1, v.y, input[v.y][v.x-1], result[v.y][v.x-1]})
			}
		}
		if v.x < len(input[0])-1 {
			if !visited[v.y][v.x+1] {
				result[v.y][v.x+1] = math.Max(input[v.y][v.x+1], math.Min(result[v.y][v.x+1], v.level))
				heap.Push(queue, Level{v.x + 1, v.y, input[v.y][v.x+1], result[v.y][v.x+1]})
			}
		}
	}

	sum := 0
	for i := range input {
		for q := range input[0] {
			sum = sum + (int)(result[i][q]-input[i][q])
		}
	}

	return sum
}

func solvePoolProblem() int {
	var width, height int
	fmt.Scan(&height, &width)
	pool := make([][]float64, height)
	for i := range pool {
		pool[i] = make([]float64, width)
		for q := range pool[0] {
			fmt.Scan(&pool[i][q])
		}
	}
	return fillPool(pool)
}

func main() {
	var numberOfCases int
	fmt.Scan(&numberOfCases)
	tests := make([]int, numberOfCases)
	for i := 0; i < numberOfCases; i++ {
		tests[i] = solvePoolProblem()
	}

	for i := 0; i < numberOfCases; i++ {
		fmt.Println(tests[i])
	}
}
