package part2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Calibration(s string) (result int64) {
	file, err := os.Open(s)
	if err != nil {
		log.Println(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value := ProcessCalibration(scanner.Text())
		result += int64(value)
	}

	return result
}

func preprocessText(s string) (result []string) {
	chars := ""
	for i, char := range s {
		_, err := strconv.Atoi(string(char))
		if err != nil {
			chars += string(char)

			if i == len(s)-1 {
				result = append(result, chars)
			}

			continue
		}

		result = append(result, chars)
		result = append(result, string(char))
		chars = ""
	}

	return result
}

func ProcessCalibration(s string) (result int) {
	result = getValueText(preprocessText(s))
	return result
}

func getValueText(raws []string) (result int) {
	var (
		tmp             []int
		mapChangeNumber = map[string]string{
			"one":   "oonee",
			"two":   "ttwoo",
			"three": "tthreee",
			"five":  "ffivee",
			"six":   "ssixx",
			"seven": "ssevenn",
			"eight": "eeightt",
			"nine":  "nninee",
			"four":  "ffourr",
			"ten":   "ttenn",
			"zero":  "zzeroo",
		}

		mapNumber = map[string]string{
			"one":   "1",
			"two":   "2",
			"three": "3",
			"five":  "5",
			"six":   "6",
			"seven": "7",
			"eight": "8",
			"nine":  "9",
			"four":  "4",
			"ten":   "10",
			"zero":  "0",
		}
	)

	for _, raw := range raws {
		val, err := strconv.Atoi(raw)
		if err == nil {
			tmp = append(tmp, val) //if number
			continue
		}

		numberString := raw
		for key := range mapChangeNumber {
			if strings.Contains(numberString, key) {
				val, found := mapChangeNumber[key]

				if found {
					numberString = strings.ReplaceAll(numberString, key, val)
				}
			}
		}

		numbers := numberString
		for key := range mapChangeNumber {
			if strings.Contains(numberString, key) {
				val, found := mapNumber[key]
				if found {
					numbers = strings.ReplaceAll(numbers, key, val)
				}
			}
		}

		for _, n := range numbers {
			val, err := strconv.Atoi(string(n))
			if err != nil {
				continue
			}

			tmp = append(tmp, val)
		}
	}

	return calculateNumbers(tmp)
}

func calculateNumbers(numbers []int) (result int) {
	switch len(numbers) {
	case 0:
		result = 0
	case 1:
		result = numbers[0]*10 + numbers[0]
	default:
		result = (numbers[0] * 10) + (numbers[len(numbers)-1])
	}

	return result
}
