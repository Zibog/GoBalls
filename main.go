package main

import (
	"fmt"
	"time"
)

func main() {
	s0 := StateFromString("YRGBWYRGBEYRGBWYRGBWYRGBW").MoveRight(2) // 1 step
	//s0 = StateFromString("WGRYBBWYRGYRGBWYRGBWYRGBE") // unsolvable
	//s0 = StateFromString("YRGBWWYRGBBGYRWEGRBYBRGYW") // unsolvable
	s0 = StateFromString("YRGBWRGBWYGBWYRBWYRGYRGBE") // 5 steps
	s0.Print()
	h := Heuristic(s0)
	fmt.Println(h)
	start := time.Now()
	//res := BFS(s0)
	//res := DFSIterative(s0)
	res := DFSRecursive(s0)
	//res := aStar(s0)
	elapsed := time.Now().Sub(start)
	fmt.Println("Board:")
	res.Print()
	fmt.Println("Time:", elapsed)
	fmt.Println("Steps:", res.Steps)

	//fmt.Println(s0.MoveUp().GetMoves().Len())
	/*sLeft := s0.MoveLeft(0)
	sLeft.Print()
	sRight := sLeft.MoveRight(0)
	sRight.Print()
	sUp := sRight.MoveUp()
	sUp.Print()
	sDown := sUp.MoveDown()
	sDown.Print()*/

	/*sP := StateFromString("YRGBWYRGBEYRGBWYRGBWYRGBW")
	s0 := sP.MoveRight(2)
	s1 := sP.MoveRight(2)
	fmt.Println(s0.Board == s1.Board)*/
}
