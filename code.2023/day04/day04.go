package day04

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func ScratchedCards(s string) int64 {
	var result int64
	file, err := os.Open(s)
	if err != nil {
		log.Println(err)
		return result
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		sum := ProcessScartchcards(scanner.Text())
		fmt.Println(sum)
		result += sum
	}

	return result
}

func ProcessScartchcards(s string) int64 {
	text := strings.Split(s, ":")
	if len(text) < 2 {
		fmt.Println("invalid input")
		return 0
	}
	card := strings.Split(strings.TrimSpace(text[1]), "|")
	if len(text) < 2 {
		fmt.Println("invalid cards")
		return 0
	}

	arrWinningNumber := stringToMap(card[0])
	arrAppearNumbers := stringToMap(card[1])

	maps := mapContains(arrWinningNumber, arrAppearNumbers)
	result := countScratchedCards(maps)

	return result
}

func stringToMap(s string) map[int]bool {
	var cards = make(map[int]bool)

	tmp := strings.Split(s, " ")
	for _, v := range tmp {
		val, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("error while convert string to int")
			continue
		}

		cards[val] = true
	}

	return cards
}

func mapContains(source, target map[int]bool) []int {
	var result []int

	for val := range source {
		_, found := target[val]
		if found {
			result = append(result, val)
		}
	}

	return result
}

func countScratchedCards(numbers []int) int64 {
	fmt.Println(numbers)
	if len(numbers) == 0 {
		return 0
	}

	if len(numbers) == 1 {
		return 1
	}

	result := math.Pow(2, float64(len(numbers)-1))
	return int64(result)
}
