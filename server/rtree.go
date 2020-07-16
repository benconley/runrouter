package main

import (
	rtreego "github.com/dhconnelly/rtreego"
	osm "github.com/paulmach/osm"
)

type RTNode struct {
	ID   osm.NodeID
	geom rtreego.Point
}

const NumDims int = 2
const MinNodes = 0
const MaxNodes = 50
const NodeTolerance = 0.001

func (node RTNode) Bounds() *rtreego.Rect {
	return node.geom.ToRect(NodeTolerance)
}

func NewTreeFromNodes(nodes map[osm.NodeID]*osm.Node) *rtreego.Rtree {
	var points []rtreego.Spatial
	for i, node := range nodes {
		rtnode := RTNode{
			ID:   i,
			geom: rtreego.Point{node.Lon, node.Lat},
		}
		points = append(points, rtnode)
	}
	tree := rtreego.NewTree(
		NumDims,
		MinNodes,
		MaxNodes,
		points...,
	)
	return tree
}
