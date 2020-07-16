package main

import (
	"context"
	"fmt"
	"os"

	osm "github.com/paulmach/osm"
	osmxml "github.com/paulmach/osm/osmxml"
)

var RoadTypes map[string]int

func init() {
	RoadTypes = make(map[string]int)
	RoadTypes["unknown"] = 0
	RoadTypes["motorway"] = 1
	RoadTypes["trunk"] = 2
	RoadTypes["primary"] = 3
	RoadTypes["secondary"] = 4
	RoadTypes["tertiary"] = 5
	RoadTypes["unclassified"] = 6
	RoadTypes["residential"] = 7
	RoadTypes["service"] = 8
	RoadTypes["motorway_link"] = 9
	RoadTypes["trunk_link"] = 10
	RoadTypes["primary_link"] = 11
	RoadTypes["secondary_link"] = 12
	RoadTypes["tertiary_link"] = 13
	RoadTypes["living_street"] = 14
	RoadTypes["footway"] = 15
	RoadTypes["path"] = 16
	RoadTypes["track"] = 17
}

func LoadGraph(osmFile string) (*Graph, error) {
	file, err := os.Open(osmFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := osmxml.New(context.Background(), file)
	defer scanner.Close()

	lookup := make(map[osm.NodeID]*osm.Node)
	edges := make(map[osm.NodeID][]Edge)
	nodeUsed := make(map[osm.NodeID]bool)

	for scanner.Scan() {
		switch e := scanner.Object().(type) {
		case *osm.Node:
			lookup[e.ID] = e
		case *osm.Way:
			tags := e.TagMap()
			roadTypeKey, ok := tags["highway"]
			if !ok {
				roadTypeKey = "unknown"
			}

			roadType, ok := RoadTypes[roadTypeKey]
			if !ok {
				fmt.Println(roadTypeKey)
				continue
			}

			nodeIDs := e.Nodes.NodeIDs()
			for i := 1; i < len(nodeIDs); i++ {
				n1 := lookup[nodeIDs[i-1]]
				n2 := lookup[nodeIDs[i]]
				nodeUsed[n1.ID] = true
				nodeUsed[n2.ID] = true
				distance := distance(n1, n2)

				edges[n1.ID] = append(edges[n1.ID], Edge{n2.ID, distance, roadType})
				edges[n2.ID] = append(edges[n2.ID], Edge{n1.ID, distance, roadType})
			}
		}
	}

	scanErr := scanner.Err()
	if scanErr != nil {
		panic(scanErr)
	}

	nodes := make(map[osm.NodeID]*osm.Node)
	for _, node := range lookup {
		if nodeUsed[node.ID] {
			nodes[node.ID] = node
		}
	}

	tree := NewTreeFromNodes(nodes)

	graph := &Graph{nodes, edges, tree}

	return graph, nil
}
