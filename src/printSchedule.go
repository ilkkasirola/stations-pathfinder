package main

import (
	"fmt"
	// "slices"
	"strings"
)

func PrintSchedule(g *GraphList, paths [][]int, source, sink string, numOfTrains int) {
	// sourceIdx := slices.Index(g.stationsNames, source)

	// maxFlow := len(paths) // How many trains can travel in parallel

	if len(paths) == 0 {
		fmt.Println("No path available from source to sink")
		return
	}

	type Train struct {
		id       int
		path     []int
		position int
	}

	var active []*Train
	nextTrainID := 1
	finished := 0
	turn := 0
	// Get the path length from the first path
	for finished < numOfTrains {
		turn++
		var output []string

		// Move existing trains forward
		var stillActive []*Train
		for _, t := range active {
			t.position++

			if t.position < len(t.path) {
				station := g.stationsNames[t.path[t.position]]
				output = append(output, fmt.Sprintf("T%d-%s", t.id, station))
				stillActive = append(stillActive, t)
			} else {
				finished++
			}
		}
		active = stillActive

		// Launch new trains (one per path per turn)
		for i := 0; i < len(paths) && nextTrainID <= numOfTrains; i++ {
			path := paths[i]

			// start at first move (skip source)
			if len(path) > 1 {
				t := &Train{
					id:       nextTrainID,
					path:     path,
					position: 1,
				}
				station := g.stationsNames[path[1]]
				output = append(output, fmt.Sprintf("T%d-%s", t.id, station))

				active = append(active, t)
				nextTrainID++
			}
		}

		if len(output) > 0 {
			fmt.Println(strings.Join(output, " "))
		}
	}
}
