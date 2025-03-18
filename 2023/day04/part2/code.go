package part1

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
	cardIndexes []int
)

func ScratchedCards(s string) int {
	var result int
	file, err := os.Open(s)
	if err != nil {
		log.Println(err)
		return result
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

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
