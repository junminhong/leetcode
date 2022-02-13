package main

import (
	"log"
	"time"
)

// kernel: find highest vertical and furthest point
func maxArea(height []int) int {
	return v3(height)
}

func v2(height []int) int {
	var (
		channel = make(chan int)
		total   = 0
	)
	for index := 0; index < len(height); index++ {
		go v2Handler(index, height, channel)
	}
	go func() {
		for totalTmp := range channel {
			if totalTmp > total {
				total = totalTmp
			}
		}
	}()
	time.Sleep(time.Hour)
	return total
}
func v2Handler(index int, height []int, channel chan int) {
	var (
		total    int = 0
		totalTmp int = 0
		pointTmp int = 1
	)
	for subIndex := index + 1; subIndex < len(height); subIndex++ {
		totalTmp = 0
		if height[index] > height[subIndex] {
			totalTmp += height[subIndex] * pointTmp
			pointTmp += 1
		} else {
			totalTmp += height[index] * pointTmp
			pointTmp += 1
		}
		if totalTmp > total {
			total = totalTmp
		}
	}
	channel <- total
}

func v1(height []int) int {
	var (
		total    int = 0
		point    [2]int
		totalTmp int = 0
		pointTmp int = 1
	)
	for index := 0; index < len(height); index++ {
		pointTmp = 1
		for subIndex := index + 1; subIndex < len(height); subIndex++ {
			totalTmp = 0
			if height[index] > height[subIndex] {
				totalTmp += height[subIndex] * pointTmp
				pointTmp += 1
			} else {
				totalTmp += height[index] * pointTmp
				pointTmp += 1
			}
			if totalTmp > total {
				total = totalTmp
				point[0] = index
				point[1] = pointTmp
			}
		}
	}
	log.Println(total, point)
	return total
}

func v3(height []int) int {
	bigTotal := 0
	for index := 0; index < len(height); index++ {
		rightChannel := make(chan int)
		leftChannel := make(chan int)
		go getRightIndex(height, index, rightChannel)
		go getLeftIndex(height, index, leftChannel)
		rightIndex := <-rightChannel
		leftIndex := <-leftChannel
		if rightTotal := (rightIndex - index) * height[index]; rightTotal > bigTotal && rightIndex != -1 {
			bigTotal = rightTotal
		}
		if leftTotal := (index - leftIndex) * height[index]; leftTotal > bigTotal && leftIndex != -1 {
			bigTotal = leftTotal
		}
	}
	return bigTotal
}

func getRightIndex(height []int, index int, channel chan int) {
	rightIndex := -1
	for checkIndex := len(height) - 1; checkIndex >= 0; checkIndex-- {
		if height[index] <= height[checkIndex] {
			rightIndex = checkIndex
			break
		}
	}
	channel <- rightIndex
}

func getLeftIndex(height []int, index int, channel chan int) {
	leftIndex := -1
	for checkIndex := 0; checkIndex < len(height); checkIndex++ {
		if height[index] <= height[checkIndex] {
			leftIndex = checkIndex
			break
		}
	}
	channel <- leftIndex
}

func v4(height []int) int {
	bigTotal := 0
	for index := 0; index < len(height); index++ {
		rightIndex := -1
		leftIndex := -1
		for checkIndex := len(height) - 1; checkIndex >= 0; checkIndex-- {
			if height[index] <= height[checkIndex] {
				rightIndex = checkIndex
				break
			}
		}
		for checkIndex := 0; checkIndex < len(height); checkIndex++ {
			if height[index] <= height[checkIndex] {
				leftIndex = checkIndex
				break
			}
		}
		if rightTotal := (rightIndex - index) * height[index]; rightTotal > bigTotal && rightIndex != -1 {
			bigTotal = rightTotal
		}
		if leftTotal := (index - leftIndex) * height[index]; leftTotal > bigTotal && leftIndex != -1 {
			bigTotal = leftTotal
		}
	}
	return bigTotal
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	log.Println(v3(height))
}
