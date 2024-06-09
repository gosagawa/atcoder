package main

import (
	"container/heap"
	"fmt"
	"math"
	"runtime/debug"
)

func main() {
	m := NewMcfGraph(5)
	m.AddEdge(0, 1, 10, 2)
	m.AddEdge(0, 2, 2, 4)
	m.AddEdge(1, 2, 6, 6)
	m.AddEdge(1, 3, 6, 2)
	m.AddEdge(2, 4, 5, 2)
	m.AddEdge(3, 2, 3, 3)
	m.AddEdge(3, 4, 8, 6)

	//fmt.Println(m.edges)
	//fmt.Println(m.edgelists)
	//	fmt.Println(m.Flow(0, 4))
	result := m.Slope(0, 4, 9)
	fmt.Println(result)
}

type McfGraph struct {
	n           int
	m           int
	edges       []McfGraphEdge
	edgelists   [][]McfGraphEdge
	prevEdgeIdx [][2]int
	dualDist    [][2]int
	vis         []bool
}

type McfGraphEdge struct {
	from int
	to   int
	cap  int
	flow int
	cost int
	idx  int
	rev  int
}

func NewMcfGraph(n int) *McfGraph {

	return &McfGraph{
		n:           n,
		edgelists:   make([][]McfGraphEdge, n),
		prevEdgeIdx: make([][2]int, n),
		dualDist:    make([][2]int, n),
		vis:         make([]bool, n),
	}
}

func (m *McfGraph) assert(cond bool) {
	if !cond {
		debug.PrintStack()
		panic("assertion failed")
	}
}

func (m *McfGraph) AddEdge(from, to, cap, cost int) int {
	m.assert(0 <= from && from < m.n)
	m.assert(0 <= to && to < m.n)
	m.assert(0 <= cap)
	m.assert(0 <= cost)

	m.edges = append(m.edges, McfGraphEdge{
		from: from,
		to:   to,
		cap:  cap,
		cost: cost,
		idx:  len(m.edgelists[from]),
		rev:  len(m.edgelists[to]),
	})
	m.m = len(m.edges)

	m.edgelists[from] = append(m.edgelists[from], McfGraphEdge{
		to:   to,
		cap:  cap,
		cost: cost,
		idx:  len(m.edgelists[from]),
		rev:  len(m.edgelists[to]),
	})
	m.edgelists[to] = append(m.edgelists[to], McfGraphEdge{
		to:   from,
		cap:  0,
		cost: -cost,
		idx:  len(m.edgelists[to]),
		rev:  len(m.edgelists[from]) - 1,
	})
	return m.m - 1
}

func (m *McfGraph) GetEdge(i int) McfGraphEdge {
	m.assert(0 <= i && i < m.m)
	return m.edges[i]
}

func (m *McfGraph) Edges() []McfGraphEdge {
	return m.edges
}

func (m *McfGraph) Flow(s, t int, arg ...int) [2]int {
	m.assert(len(arg) <= 1)
	flowLimit := math.MaxInt
	if len(arg) == 1 {
		flowLimit = arg[0]
	}
	slope := m.Slope(s, t, flowLimit)

	return slope[len(slope)-1]
}

func (m *McfGraph) Slope(s, t int, arg ...int) [][2]int {
	m.assert(len(arg) <= 1)
	flowLimit := math.MaxInt
	if len(arg) == 1 {
		flowLimit = arg[0]
	}

	m.assert(0 <= s && s < m.n)
	m.assert(0 <= t && t < m.n)
	m.assert(s != t)

	var flow, cost int
	prevCostPerFlow := -1
	result := make([][2]int, 0)
	for {

		/*
			for _, l := range m.edgelists {
				ts := make([]string, len(l))
				for j, v := range l {
					ts[j] = fmt.Sprintf("{to:%v cap:%v}", v.to, v.cap)
				}
				fmt.Println(strings.Join(ts, ""))
			}
			fmt.Println(m.prevEdgeIdx)
			fmt.Println("")
		*/

		if flow >= flowLimit {
			break
		}

		if !m.dualRef(s, t) {
			break
		}
		c := flowLimit - flow
		for v := t; v != s; v = m.prevEdgeIdx[v][0] {
			pidx := m.prevEdgeIdx[v][1]
			from := m.edgelists[v][pidx].to
			fidx := m.edgelists[v][pidx].rev

			if c > m.edgelists[from][fidx].cap {
				c = m.edgelists[from][fidx].cap
			}
		}
		for v := t; v != s; v = m.prevEdgeIdx[v][0] {
			pidx := m.prevEdgeIdx[v][1]
			from := m.edgelists[v][pidx].to
			fidx := m.edgelists[v][pidx].rev

			m.edgelists[v][pidx].cap += c
			m.edgelists[from][fidx].cap -= c
		}
		d := -m.dualDist[s][0]
		flow += c
		cost += c * d
		if prevCostPerFlow == d {
			result = result[:len(result)-1]
		}
		result = append(result, [2]int{flow, cost})
		prevCostPerFlow = d
	}

	for i := 0; i < m.m; i++ {
		e := m.edgelists[m.edges[i].from][m.edges[i].idx]
		m.edges[i].flow = m.edges[i].cap - e.cap
	}

	return result
}

type McfGraphQ struct {
	key int
	to  int
}

type McfGraphPQ []*McfGraphQ

func (pq McfGraphPQ) Len() int { return len(pq) }

func (pq McfGraphPQ) Less(i, j int) bool {
	return pq[i].key < pq[j].key
}

func (pq McfGraphPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *McfGraphPQ) Push(x any) {
	item := x.(*McfGraphQ)
	*pq = append(*pq, item)
}

func (pq *McfGraphPQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func (m *McfGraph) dualRef(s, t int) bool {
	for i := 0; i < m.n; i++ {
		m.vis[i] = false
		m.dualDist[i][1] = math.MaxInt32
	}

	queMin := []int{}
	que := new(McfGraphPQ)
	heap.Init(que)

	m.dualDist[s][1] = 0
	queMin = append(queMin, s)
	for {
		if len(queMin) == 0 && que.Len() == 0 {
			break
		}
		v := 0
		if len(queMin) > 0 {
			v = queMin[len(queMin)-1]
			queMin = queMin[:len(queMin)-1]
		} else {
			v = heap.Pop(que).(*McfGraphQ).to
		}
		if m.vis[v] {
			continue
		}
		m.vis[v] = true
		if v == t {
			break
		}

		dualV := m.dualDist[v][0]
		distV := m.dualDist[v][1]
		for _, e := range m.edgelists[v] {
			if e.cap == 0 {
				continue
			}

			cost := e.cost - m.dualDist[e.to][0] + dualV
			if m.dualDist[e.to][1]-distV > cost {
				distTo := distV + cost
				m.dualDist[e.to][1] = distTo
				m.prevEdgeIdx[e.to] = [2]int{v, e.rev}
				if distTo == distV {
					queMin = append(queMin, e.to)
				} else {
					heap.Push(que, &McfGraphQ{distTo, e.to})
				}
			}
		}
	}
	if !m.vis[t] {
		return false
	}

	for v := 0; v < m.n; v++ {
		if !m.vis[v] {
			continue
		}
		// dual[v] = dual[v] - dist[t] + dist[v]
		//         = dual[v] - (shortest(s, t) + dual[s] - dual[t]) +
		//         (shortest(s, v) + dual[s] - dual[v]) = - shortest(s,
		//         t) + dual[t] + shortest(s, v) = shortest(s, v) -
		//         shortest(s, t) >= 0 - (n-1)C
		m.dualDist[v][0] -= m.dualDist[t][1] - m.dualDist[v][1]
	}
	return true
}
