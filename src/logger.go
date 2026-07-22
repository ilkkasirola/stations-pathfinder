package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	PrefixError	string = `[ERROR] `
	PrefixLog	string = `[LOG]   `
)

var (
	Usage						string = "go run . [path to file containing network map] [start station] [end station] [number of trains]"
	ErrArgsCount 				string = "Incorrect number of command line arguments"
	ErrStartStationNotExist 	string = "Start station does not exist"
	ErrEndStationNotExist		string = "End station does not exist"
	ErrStationsSame				string = "Start and end station are the same"
	ErrNoPath					string = "No path between the start and end stations"
	ErrDuplicateConnections		string = "Duplicate connections: %s"
	ErrTrainsCount				string = "Number of trains is not a valid positive integer"
	ErrInvalidCoordinates		string = "Coordinates are not valid positive integers"
	ErrDuplicateCoordinates		string = "Two stations exist at the exact same coordinate location: %s"
	ErrConnectionNotExist		string = "Connection does not exist: %s"
	ErrDuplicateStations		string = "Duplicate station names: %s"
	ErrInvalidStationNameArg	string = "Invalid station name in argument"
	ErrNoStations				string = "Map does not contain stations section"
	ErrNoConnections			string = "Map does not contain connections section"
	ErrTooManyStations			string = "Map has over 10 000 stations"
	ErrMalformedStation			string = "Malformed station data: %s"
	ErrMalformedConnection		string = "Malformed connection data: %s"
	ErrStationNotExist			string = "Station does not exist: %s"
)

func PrintErr(messages ...string) {
	message := strings.Join(messages, " ")
	message = fmt.Sprintf("%s%s\n", PrefixError, message)
	os.Stderr.Write([]byte(message))
}

// Prints one of the defined messages with string format

// Example:
// PrintErrArgs(ErrDuplicateStation, "hakaniemi", "kalasatama")
// Will print out:
// [ERROR] Duplicate station names: [hakaniemi kalasatama]
func PrintErrArgs(errMessage string, args ...string) {
	message := fmt.Sprintf("%s%s\n", PrefixError, errMessage)
	message = fmt.Sprintf(message, args)
	os.Stderr.Write([]byte(message))
}

func Log(messages ...string) {
	message := strings.Join(messages, " ")
	fmt.Println(fmt.Sprintf("%s%s", PrefixLog, message))
}
