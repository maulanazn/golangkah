package main

import "fmt"

type Score struct {
	scoreId string
	Score   []ScoreDetail
}

type ScoreDetail struct {
	id    string
	name  string
	value string
}

func main() {
	var maulanaScore ScoreDetail = ScoreDetail{"1a2b738m-dsffd-9883n", "maulana", "A"}

	var allScore []ScoreDetail = []ScoreDetail{maulanaScore}
	var scoreInit Score = Score{"1723-ndnm2-233m", allScore}

	fmt.Printf("All Score %q", scoreInit)
	fmt.Printf("Maulana Score %q", maulanaScore)

}
