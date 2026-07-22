package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	maxStations int = 10_000
)

var graphList *GraphList

func ParseMap(input string) (*GraphList, bool) {
	graphList = NewGraphlist()

	var parseStations bool = false
	var parseConnections bool = false

	if !strings.Contains(input, "stations:") {
		PrintErr(ErrNoStations)
		return graphList, false
	}

	if !strings.Contains(input, "connections:") {
		PrintErr(ErrNoConnections)
		return graphList, false
	}

	for line := range strings.SplitSeq(input, "\n") {
		var commentIndex int = strings.Index(line, "#")
		if commentIndex != -1 {
			line = line[:commentIndex]
		}

		line = strings.ReplaceAll(line, " ", "")
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "stations:") {
			parseStations = true
			continue
		}

		if strings.HasPrefix(line, "connections:") {
			parseStations = false
			parseConnections = true
			continue
		}

		if parseStations {
			station, err := parseStation(line)
			if err != nil {
				return graphList, false
			}
			ok := graphList.AddStation(&station)
			if !ok {
				return graphList, false
			}
		}
		if parseConnections {
			begin, end, err := parseConnection(line)
			if err != nil {
				return graphList, false
			}
			ok := graphList.AddConnection(begin, end)
			if !ok {
				return graphList, false
			}
		}
	}

	if len(graphList.stations) > maxStations {
		PrintErr(ErrTooManyStations)
		return graphList, false
	}

	return graphList, true
}

func parseStation(input string) (Station, error) {
	var stationValues []string = strings.Split(input, ",")
	if len(stationValues) != 3 {
		PrintErrArgs(ErrMalformedStation, input)
		return Station{}, errors.New(fmt.Sprintf(ErrMalformedStation, input))
	}

	name := stationValues[0]
	x, err := strconv.Atoi(stationValues[1])
	if err != nil || x < 0 {
		PrintErr(ErrInvalidCoordinates)
		return Station{}, errors.New(ErrInvalidCoordinates)
	}
	y, err := strconv.Atoi(stationValues[2])
	if err != nil || y < 0 {
		PrintErr(ErrInvalidCoordinates)
		return Station{}, errors.New(ErrInvalidCoordinates)
	}

	return Station{
		Name: name,
		X:    x,
		Y:    y,
	}, nil
}

func parseConnection(input string) (string, string, error) {
	var connectionValues []string = strings.Split(input, "-")
	if len(connectionValues) != 2 {
		PrintErr(ErrMalformedConnection, input)
		return "", "", errors.New(fmt.Sprintf(ErrMalformedConnection, input))
	}

	return connectionValues[0], connectionValues[1], nil
}
