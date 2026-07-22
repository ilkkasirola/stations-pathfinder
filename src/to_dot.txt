package main


import (
	"fmt"
	"strings"
)

const (
	shape string = "ellipse"
	fixedsize bool = false
	prefix string = "graph stations \n{rankdir=LR;"
	suffix string = "\n}"
)


func ToDot(stations []Station, connections []Connection) string {
	var b strings.Builder
	b.WriteString(prefix)
	fmt.Fprintf(&b, "\n\nnode [shape=%s, fixedsize=%t];\n\n", shape, fixedsize)

	for _, station := range stations {
		fmt.Fprintf(&b, `"%s";`, station.Name)
		fmt.Fprintf(&b, "\n")
	}
	b.WriteString("\n")
	for _, connection := range connections {
		fmt.Fprintf(&b, `"%s" -- "%s";`, connection.Begin, connection.End)
		fmt.Fprintf(&b, "\n")
	}
	b.WriteString(suffix)

	return b.String()
}
