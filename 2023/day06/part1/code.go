package part1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Race(s string) int {
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	var raws []string

	for scanner.Scan() {
		raws = append(raws, scanner.Text())
	}

	return processRace(raws)
}

func processRace(raws []string) int {
	var (
		sum       = 1
		times     []int
		distances []int
	)

	for i, raw := range raws {
		s := strings.Split(strings.TrimSpace(strings.Split(raw, ":")[1]), " ")
		if i == 0 {
			times = append(times, toInt(s)...)
		}
		if i == 1 {
			distances = append(distances, toInt(s)...)
		}
	}

	for i := 0; i < len(times); i++ {
		attempt := CalcRac(times[i], distances[i])
		if attempt == 0 {
			continue
		}
		sum *= attempt
	}

	return sum
}

func CalcRac(times, distance int) int {
	var result int

	for i := 0; i < times; i++ {
		currentMaxDistance := i * (times - i)
		if currentMaxDistance > distance {
			result += 1
		}
	}

	return result
}

func toInt(s []string) (res []int) {
	for _, v := range s {
		i, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		res = append(res, i)
	}

	return
}
