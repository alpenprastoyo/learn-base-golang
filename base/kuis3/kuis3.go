package main

import "fmt"

func main() {
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var goodScores []int
	var average float32
	var totalScore float32
	totalScore = 0

	for _, value := range scores {
		totalScore += float32(value)

		if value >= 90 {
			goodScores = append(goodScores, value)
		}

	}

	average = totalScore / float32(len(scores))

	fmt.Println("The average of scores is :", average)
	fmt.Println("The good scores is :", goodScores)

}
