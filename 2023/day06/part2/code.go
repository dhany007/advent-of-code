package part2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Race(s string) int64 {
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

func processRace(raws []string) int64 {
	var (
		rawTimes     string
		rawDistances string
	)

	for i, raw := range raws {
		tmp := strings.Split(strings.TrimSpace(strings.Split(raw, ":")[1]), " ")

		if i == 0 {
			rawTimes += strings.Join(tmp, "")
		}

		if i == 1 {
			rawDistances += strings.Join(tmp, "")
		}
	}

	times, err := strconv.ParseInt(rawTimes, 10, 64)
	if err != nil {
		return 0
	}

	distances, err := strconv.ParseInt(rawDistances, 10, 64)
	if err != nil {
		return 0
	}

	return calculateRace(times, distances)
}

func calculateRace(times, distance int64) int64 {
	var result int64

	var i int64
	for i = 0; i < times; i++ {
		currentMaxDistance := i * (times - i)
		if currentMaxDistance > distance {
			result += 1
		}
	}

	return result
}
