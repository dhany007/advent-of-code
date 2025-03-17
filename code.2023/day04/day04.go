package day04

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type (
	Card struct {
		Original int
		Copy     int
		Count    int
	}
)

var (
	cards       = make(map[int]Card)
	cardIndexes = []int{}
)

func ScratchedCards(s string) int {
	var result int
	file, err := os.Open(s)
	if err != nil {
		log.Println(err)
		return result
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		count := getCardCount(scanner.Text())
		preProcessCard(count, i)
		i += 1
	}
	result = processCard()

	return result
}

func preProcessCard(count int, key int) {
	cardIndexes = append(cardIndexes, key)
	cards[key] = Card{
		Original: 1,
		Count:    count,
	}
}

func processCard() (total int) {
	for _, key := range cardIndexes {
		card, found := cards[key]
		if !found {
			continue
		}

		var incrementCopy []int
		for i := key + 1; i < card.Count+key+1; i++ {
			tmp, found := cards[i]
			if found {
				cards[i] = Card{
					Original: tmp.Original,
					Count:    tmp.Count,
					Copy:     tmp.Copy + 1,
				}
			}
			incrementCopy = append(incrementCopy, i)
		}

		if card.Copy > 0 {
			for i := 0; i < card.Copy; i++ {
				for _, val := range incrementCopy {
					tmp, found := cards[val]
					if found {
						cards[val] = Card{
							Original: tmp.Original,
							Count:    tmp.Count,
							Copy:     tmp.Copy + 1,
						}
					}
				}
			}
		}
	}

	for _, value := range cards {
		total = total + value.Original + value.Copy
	}

	return
}

func getCardCount(c string) int {
	text := strings.Split(c, ":")
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

	count := mapContains(arrWinningNumber, arrAppearNumbers)

	return count
}

func stringToMap(s string) map[int]bool {
	var cards = make(map[int]bool)

	tmp := strings.Split(s, " ")
	for _, v := range tmp {
		val, err := strconv.Atoi(v)
		if err != nil {
			continue
		}

		cards[val] = true
	}

	return cards
}

func mapContains(source, target map[int]bool) (result int) {
	for val := range source {
		_, found := target[val]
		if found {
			result += 1
		}
	}

	return result
}

/*
below solution of day 4 part 1

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

*/
