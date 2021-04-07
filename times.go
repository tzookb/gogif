package main

// https://github.com/urfave/cli/blob/master/docs/v2/manual.md#getting-started

import (
	"math"

	"github.com/uniplaces/carbon"
)

func getThen() (*carbon.Carbon, error) {
	then, err := carbon.Create(2021, 4, 4, 18, 0, 0, 0, "Local")
	// then, err := carbon.CreateFromDate(2021, time.April, 1, "Local")
	if err != nil {
		return nil, err
	}
	then.SetHour(0)
	then.SetMinute(0)
	then.SetSecond(0)
	return then, nil
}

type TimeBlock struct {
	days    int
	hours   int
	minutes int
	seconds int
}

func getDiff(from *carbon.Carbon, to *carbon.Carbon) TimeBlock {
	days := int(from.DiffInDays(to, true))
	hours := int(math.Floor(float64(from.DiffInHours(to, true) % 24)))
	minutes := int(math.Floor(float64(from.DiffInMinutes(to, true) % 60)))
	seconds := int(math.Floor(float64(from.DiffInSeconds(to, true) % 60)))

	return TimeBlock{days, hours, minutes, seconds}
}

func getDiffsBack(count int, from *carbon.Carbon, to *carbon.Carbon) *[]TimeBlock {
	timeBlocks := []TimeBlock{}

	i := 0
	for i < count {
		tb := getDiff(from, to)
		timeBlocks = append(timeBlocks, tb)
		from = from.SubSecond()
		i++
	}

	return &timeBlocks
}
