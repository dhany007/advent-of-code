package day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	labelRed   string = "red"
	labelGreen string = "green"
	labelBlue  string = "blue"

	maxRed   int = 12
	maxGreen int = 13
	maxBlue  int = 14
)

func SumValidBag(s string) (sum int) {
	file, err := os.Open(s)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	increment := 1
	for scanner.Scan() {
		a := PreprocessText(scanner.Text())
		if a {
			sum += increment
		}
		increment += 1
	}

	return
}

func PreprocessText(text string) bool {
	raw := strings.Split(text, ":")

	if len(raw) < 2 {
		fmt.Println("invalid length string")
		return false
	}

	cubes := strings.Split(raw[1], ";")
	for _, cube := range cubes {
		tmpCubes := strings.Split(strings.TrimSpace(cube), ", ")
		if !IsValidCubeAttempt(tmpCubes) {
			return false
		}
	}

	return true
}

func IsValidCubeAttempt(cubes []string) bool {
	if len(cubes) == 0 {
		fmt.Println("invalid length cubes")
		return false
	}

	for _, s := range cubes {
		s = strings.TrimSpace(s)

		red := stringCubeToInt(s, labelRed)
		green := stringCubeToInt(s, labelGreen)
		blue := stringCubeToInt(s, labelBlue)

		if !isValidRule(red, green, blue) {
			return false
		}
	}

	return true
}

func stringCubeToInt(cube, label string) (count int) {
	if strings.Contains(cube, label) {
		val, _ := strconv.Atoi(strings.Split(cube, " ")[0])
		count += val
	}

	return count
}

func isValidRule(red, green, blue int) bool {
	return red <= maxRed && green <= maxGreen && blue <= maxBlue
}
