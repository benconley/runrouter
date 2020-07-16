package main

import (
	"container/heap"
	"fmt"
	"math"
	"strings"

	rtreego "github.com/dhconnelly/rtreego"
	osm "github.com/paulmach/osm"
)

func distance(a, b *osm.Node) float64 {
	const diameter = 2 * EarthRadiusMeters
	lat1 := Radians(a.Lat)
	lat2 := Radians(b.Lat)
	latH := math.Sin((lat1 - lat2) / 2)
	latH *= latH
	lonH := math.Sin(Radians(a.Lon-b.Lon) / 2)
	lonH *= lonH
	tmp := latH + math.Cos(lat1)*math.Cos(lat2)*lonH
	return diameter * math.Asin(math.Sqrt(tmp))
}

type Edge struct {
	ID       osm.NodeID
	Distance float64
	RoadType int
}

type Graph struct {
	Nodes map[osm.NodeID]*osm.Node
	Edges map[osm.NodeID][]Edge
	Tree  *rtreego.Rtree
}

func (graph *Graph) Search(src *osm.Node, maxDistance float64) []SearchResult {
	var queue PriorityQueue
	routes := make([]SearchResult, NumRoutes)
	i := 0

	heap.Push(&queue, NewItem(src))
	for len(queue) > 0 {
		item := heap.Pop(&queue).(*Item)

		if item.Distance > maxDistance*1.02 {
			continue
		}
		if item.Distance > maxDistance*0.98 {
			routes[i] = graph.CreateSearchResult(item)
			i++

			if i >= NumRoutes {
				break
			}
		}
		for _, edge := range graph.Edges[item.ID] {
			newItem := item.Follow(edge)
			if item.Next != nil && item.Next.ID == newItem.ID {
				continue
			}
			heap.Push(&queue, newItem)
		}
	}
	return routes
}

type SearchResult struct {
	Nodes    []*osm.Node
	Distance float64
}

func (s SearchResult) String() string {
	str := make([]string, 0)

	for _, node := range s.Nodes {
		str = append(str, fmt.Sprintf("[%f,%f]", node.Lon, node.Lat))
	}

	return "[" + strings.Join(str, ",") + "]"
}

func (graph *Graph) CreateSearchResult(item *Item) SearchResult {
	path := make([]*osm.Node, item.Depth+1)
	distance := item.Distance
	for item != nil {
		path[item.Depth] = graph.Nodes[item.ID]
		item = item.Next
	}
	return SearchResult{path, distance}
}
