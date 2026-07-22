package main

import (
	"container/list"
	"strings"
)


// big fucking search
func Bfs(g *GraphList, start, end string) []string {
	q := list.New()
	q.PushBack(g.stations[start].Name)

	parents := make(map[string]string)
	parents[start] = ""

	for q.Len() > 0 {
		currentStation := q.Front().Value.(string)
		q.Remove(q.Front())

		if strings.EqualFold(currentStation, end) {
			var route []string
			for len(currentStation) > 0 {
				route = append([]string{currentStation}, route...)
				currentStation = parents[currentStation]
			}

			return route
		}

		for _, neighbor := range g.adjacentList[currentStation] {
			if _, visited := parents[neighbor]; !visited {
				parents[neighbor] = currentStation
				q.PushBack(g.stations[neighbor].Name)
			}
		}
	}

	return []string{"invalid"}
}
