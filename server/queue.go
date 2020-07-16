package main

import osm "github.com/paulmach/osm"

type Item struct {
	ID            osm.NodeID
	Distance      float64
	TotalDistance float64
	Score         int
	Depth         int
	Next          *Item
}

func NewItem(node *osm.Node) *Item {
	return &Item{ID: node.ID}
}

func (item *Item) Follow(edge Edge) *Item {
	result := Item{}
	result.ID = edge.ID
	result.Distance = item.Distance + edge.Distance
	result.Score = edge.RoadType
	result.Depth = item.Depth + 1
	result.Next = item
	return &result
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].Score != pq[j].Score {
		return pq[i].Score > pq[j].Score
	} else {
		return pq[i].Distance > pq[j].Distance
	}
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}
