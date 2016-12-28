package main

import (
	"fmt"
)

type Rating struct {
	clarity     int
	originality int
	difficulty  int
}

func main() {
	a, b := read()
	aScore, bScore := calcScores(a, b)
	fmt.Println(aScore, bScore)
}

func calcScores(a, b Rating) (aScore int, bScore int) {
	aClarity, bClarity := calcCategory(a.clarity, b.clarity)
	aOriginality, bOriginality := calcCategory(a.originality, b.originality)
	aDifficulty, bDifficulty := calcCategory(a.difficulty, b.difficulty)

	aScore = aClarity + aOriginality + aDifficulty
	bScore = bClarity + bOriginality + bDifficulty

	return
}

func calcCategory(aV, bV int) (aS int, bS int) {
	if aV > bV {
		aS = 1
	}

	if bV > aV {
		bS = 1
	}

	return
}

func read() (Rating, Rating) {
	a := readRating()
	b := readRating()

	return a, b
}

func readRating() (r Rating) {
	readInts(&r.clarity, &r.originality, &r.difficulty)
	return
}

func readInts(clarity, originality, difficulty *int) {
	readInt(clarity)
	readInt(originality)
	readInt(difficulty)
}

func readInt(x *int) {
	if _, err := fmt.Scan(x); err != nil {
		panic(err)
	}
}
