package main

import (
	"github.com/gammazero/deque"
	"github.com/jupp0r/go-priority-queue"
	"math"
)

func IsSolution(s *State) bool {
	for i := 0; i < size; i++ {
		el := s.Board[0][i]
		start := 1
		if el == toi('E') {
			el = s.Board[1][i]
			start++
		}
		for j := start; j < size; j++ {
			if el != s.Board[j][i] && s.Board[j][i] != toi('E') {
				return false
			}
		}
	}
	return true
}

func BFS(s *State) *State {
	q := deque.New()
	used := make(map[[size][size]int]bool)
	q.PushBack(s)
	s.Steps = 0
	used[s.Board] = true

	for q.Len() > 0 {
		currentState := q.PopFront().(*State)

		if IsSolution(currentState) {
			return currentState
		}

		moves := currentState.GetMoves()
		for e := moves.Front(); e != nil; e = e.Next() {
			child := e.Value.(*State)
			if _, visited := used[child.Board]; !visited {
				q.PushBack(child)
			}
			used[child.Board] = true
		}
	}
	return nil
}

func DFSIterative(s *State) *State {
	q := deque.New()
	used := make(map[[size][size]int]bool)
	q.PushBack(s)
	s.Steps = 0
	used[s.Board] = true

	for q.Len() > 0 {
		currentState := q.PopBack().(*State)

		if IsSolution(currentState) {
			return currentState
		}

		moves := currentState.GetMoves()
		for e := moves.Front(); e != nil; e = e.Next() {
			child := e.Value.(*State)
			if _, visited := used[child.Board]; !visited {
				q.PushBack(child)
			}
			used[child.Board] = true
		}
	}
	return nil
}

var used = make(map[[size][size]int]int)

func DFSRecursive(s *State) *State {
	if IsSolution(s) || s.Steps > size * size {
		return s
	}
	res := State{
		Board:  [5][5]int{},
		Zero:   -1,
		Steps:  math.MaxInt,
		Parent: nil,
	}
	used[s.Board] = s.Steps
	moves := s.GetMoves()
	for e := moves.Front(); e != nil; e = e.Next() {
		child := e.Value.(*State)
		if _, visited := used[child.Board]; !visited {
			tmp := DFSRecursive(child)
			if tmp.Steps < res.Steps {
				res = *tmp
			}
		}
	}
	return &res
}

func Heuristic(s *State) float64 {
	sumR := 0
	sumG := 0
	sumB := 0
	sumY := 0
	sumW := 0
	for i := 0; i < size; i++ {
		el := s.Board[0][i]
		switch el {
		case 1:
			sumY++
		case 2:
			sumR++
		case 3:
			sumG++
		case 4:
			sumB++
		case 5:
			sumW++
		}
		for j := 1; j < size; j++ {
			if el == s.Board[j][i] || s.Board[j][i] == toi('E') {
				switch el {
				case 1:
					sumY++
				case 2:
					sumR++
				case 3:
					sumG++
				case 4:
					sumB++
				case 5:
					sumW++
				}
			}
		}
	}
	return -float64(sumR + sumG + sumB + sumY + sumW)
}

func aStar(s *State) *State {
	q := pq.New()
	used := make(map[[size][size]int]bool)
	s.Steps = 0
	currentPriority := Heuristic(s)
	q.Insert(s, currentPriority)
	used[s.Board] = true

	for q.Len() > 0 {
		st, _ := q.Pop()
		currentState := st.(*State)

		if IsSolution(currentState) {
			return currentState
		}

		used[currentState.Board] = true

		moves := currentState.GetMoves()
		for e := moves.Front(); e != nil; e = e.Next() {
			child := e.Value.(*State)
			if used[child.Board] {
				continue
			}
			currentPriority = Heuristic(child)
			q.Insert(child, currentPriority+float64(child.Steps))
		}
	}
	return nil
}