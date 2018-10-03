package Interfaces

import (
	"sort"
)

type PriorityQueue []HasSpeed

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].GetSpeed() > pq[j].GetSpeed()
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(obj HasSpeed) {
	*pq = append(*pq, obj)
}

func (pq *PriorityQueue) Pop() HasSpeed {
	sort.Sort(pq)
	item := (*pq)[0]
	*pq = (*pq)[1:]
	return item
}
