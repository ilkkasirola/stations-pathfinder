package main

import (
	"os"
	"strconv"
)


func ParseArgs() (mapfile, start, end string, trainCount int, ok bool) {
	ok = false

	if len(os.Args) != 5 {
		PrintErr(ErrArgsCount)
		Log(Usage)
		return
	}

	mapfile = os.Args[1]
	start = os.Args[2]
	end = os.Args[3]
	trainCountStr := os.Args[4]
	
	Log("Using", mapfile, "as input.")

	var err error
	trainCount, err = strconv.Atoi(trainCountStr)
	if err != nil || trainCount < 1 {
		PrintErr(ErrTrainsCount)
		return
	}

	if start == end {
		PrintErr(ErrStationsSame)
		return
	}

	ok = true
	return
}

func ReadMapFile(filename string) string {
	contents, err := os.ReadFile(filename)
	if err != nil {
		PrintErr(err.Error())
		return ""
	}
	return string(contents)
}
