package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func log(m interface{}) {
	fmt.Println(m)
}

func main() {
	// heutristic find the word with the most commonalities
	// reasoning: they want to make it hard to eliminate characters
	log("enter words")
	words := getInitialInput()

	for i := 0; i < 4; i++ {
		if len(words) == 0 {
			log("no more words. Invalid")
			return
		}
		if len(words) == 1 {
			log("the answer is: " + words[0])
			return
		}

		log(words)
		allSims := buildAllSims(words)
		guessWord := guess(allSims)
		log("guess: " + guessWord + "? enter score [0-4]")
		score := getInput()
		words = prune(allSims, guessWord, score)
	}
}

type Sims map[string]int
type SimsMap map[string]Sims

func (sm SimsMap) List() (list []string) {
	for word, _ := range sm {
		list = append(list, word)
	}
	return
}

// given a word and score,
func prune(allSims SimsMap, guess string, guessScore int) (list []string) {
	sims := allSims[guess]
	for word, score := range sims {
		if score == guessScore {
			list = append(list, word)
		}
	}
	return
}

// find the lowest max score
func guess(allSims SimsMap) string {
	// score = max repeats
	scores := map[string]int{}
	for word, sims := range allSims {
		scores[word] = maxRepeats(sims)
	}

	// find lowest max score
	var min int
	var answer string
	for word, score := range scores {
		if min == 0 || score < min {
			min = score
			answer = word
		}
	}
	return answer
}

func sim(w1 string, w2 string) int {
	sim := 0
	for i, _ := range w1 {
		log(w1)
		log(w2)
		if w1[i] == w2[i] {
			sim++
		}
	}
	return sim
}

func buildAllSims(words []string) SimsMap {
	allSims := SimsMap{}
	for _, word1 := range words {
		sims := Sims{}
		for _, word2 := range words {
			if word1 != word2 {
				sims[word2] = sim(word1, word2)
			}
		}
		allSims[word1] = sims
	}
	return allSims
}

func filter(words []string, word string, score int) []string {
	return []string{}
}

func maxRepeats(sims Sims) int {
	counts := map[int]int{}
	for _, i := range sims {
		counts[i] = counts[i] + 1
	}

	max := 0
	for _, count := range counts {
		if count > max {
			max = count
		}
	}
	return max
}

// each sims represents all possible states if that word is chosen
// we want to guaruantee to eliminate the most number of words,
// and the max number of repeats represents the worst case
// so we want to guess the word with the lowest worst case

func getInitialInput() (words []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	return strings.Split(input, " ")
}

func getInput() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for {
		input := scanner.Text()
		switch input {
		case "0", "1", "2", "3", "4":
			res, _ := strconv.Atoi(input)
			return res
		default:
			log("invalid input. enter score [1-4]")
		}
	}
}

func toString(i int) string {
	return strconv.Itoa(i)
}
