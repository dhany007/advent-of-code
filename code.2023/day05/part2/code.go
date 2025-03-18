package part2

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	rawData = make(map[int32][]string)

	seedsLabel            = "seeds"
	seedToSoil            = "seed-to-soil"
	soilToFertilizer      = "soil-to-fertilizer"
	fertilizerToWater     = "fertilizer-to-water"
	waterToLight          = "water-to-light"
	lightToTemperature    = "light-to-temperature"
	temperatureToHumidity = "temperature-to-humidity"
	humidityToLocation    = "humidity-to-location"

	fertilizerKeys = map[string]int{
		seedsLabel:            0,
		seedToSoil:            1,
		soilToFertilizer:      2,
		fertilizerToWater:     3,
		waterToLight:          4,
		lightToTemperature:    5,
		temperatureToHumidity: 6,
		humidityToLocation:    7,
	}

	maxGardener = 8

	valueSeed []int64
	incSeed   []int64
)

func GetLocationGarden(filename string) int64 {
	file, err := os.Open(filename)
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
	currentKey := ""

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" || text == " " {
			currentKey = ""
			continue
		}

		currentKey = constructData(text, currentKey)
	}

	constructSeeds()

	return processSeed()
}

func calculateCorresponds(initial int64, values []string) int64 {
	result := initial

	for _, val := range values {
		value := strings.Split(val, " ")

		currRange, _ := strconv.ParseInt(value[2], 10, 64)
		source, _ := strconv.ParseInt(value[1], 10, 64)
		target, _ := strconv.ParseInt(value[0], 10, 64)

		if source <= initial && (currRange+source) >= initial {
			dif := target - source
			result = initial + dif
			break
		}
	}

	return result
}

func constructData(t string, currentKey string) string {
	if strings.Contains(t, seedsLabel) {
		tmp := strings.TrimSpace(strings.Split(t, ":")[1])
		rawData[0] = append(rawData[0], tmp)
		return ""
	}

	tmpKey := strings.Split(t, " ")[0]
	_, found := fertilizerKeys[tmpKey]
	if found {
		return tmpKey
	}

	upsertData(t, currentKey)

	return currentKey
}

func constructSeeds() {
	seed, found := rawData[0]
	if found {
		val := strings.Split(seed[0], " ")
		for i := 0; i < len(val)-1; i++ {
			key, _ := strconv.ParseInt(val[i], 10, 64)
			inc, _ := strconv.ParseInt(val[i+1], 10, 64)
			if i%2 == 0 {
				valueSeed = append(valueSeed, key)
				incSeed = append(incSeed, inc)
			}
		}
	}
}

func processSeed() int64 {
	var minimum int64 = math.MaxInt64
	for i := 0; i < len(valueSeed); i++ {
		seed := valueSeed[i]
		for j := int64(0); j < incSeed[i]; j++ {
			result := processCorresponds(seed)
			if result < minimum {
				minimum = result
			}

			seed += 1
		}
	}

	return minimum
}

func processCorresponds(seed int64) int64 {
	corresponds := seed
	for i := 1; i < maxGardener; i++ {
		values, found := rawData[int32(i)]
		if !found {
			continue
		}

		corresponds = calculateCorresponds(corresponds, values)
	}

	return corresponds
}

func upsertData(t, key string) {
	k, ok := fertilizerKeys[key]
	if ok {
		val, found := rawData[int32(k)]
		if found {
			val = append(val, t)
			rawData[int32(k)] = val
		} else {
			rawData[int32(k)] = []string{t}
		}
	}
}
