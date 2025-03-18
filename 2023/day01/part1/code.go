package part1

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

func ProcessCalibration(s string) (result int) {
	var tmp []int

	for _, char := range s {
		val, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		}

		tmp = append(tmp, val)
	}

	switch len(tmp) {
	case 0:
		result = 0
	case 1:
		result = tmp[0]*10 + tmp[0]
	default:
		result = (tmp[0] * 10) + (tmp[len(tmp)-1])
	}

	return result
}
