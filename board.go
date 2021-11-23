package main

import (
	"container/list"
	"fmt"
)

const size = 5

type State struct {
	Board [size][size]int
	Zero int
	Steps int
	Parent *State
	//LastMove Move
}
type Move struct {
	Disk int
	// 0 == Left; 1 == Right; 2 == Up; 3 == Down
	Direction int
}

func StateFromString(board string) *State {
	if len(board) != 25 {
		panic("Wrong literal argument")
	}
	newBoard := [size][size]int{}
	zero := -1
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			temp := toi(board[i*5+j])
			newBoard[i][j] = temp
			if temp == 0 {
				zero = i*5+j
			}
		}
	}
	if zero == -1 {
		panic("Wrong literal argument")
	}
	return &State{Board: newBoard, Zero: zero, Steps: 0, Parent: nil}
}

func (this State) MoveLeft(disk int) *State {
	sNew := this
	sNew.Steps++
	sNew.Parent = &this
	/*sNew.LastMove = Move{
		Disk:      disk,
		Direction: 0,
	}*/
	temp := sNew.Board[disk][0]
	for i := 0; i < size - 1; i++ {
		sNew.Board[disk][i] = sNew.Board[disk][i + 1]
	}
	sNew.Board[disk][size - 1] = temp
	zeroPos := sNew.Zero
	if zeroPos / size == disk {
		if zeroPos % size == 0 {
			sNew.Zero = zeroPos + size - 1
		} else {
			sNew.Zero--
		}
	}
	return &sNew
}

func (this State) MoveRight(disk int) *State {
	sNew := this
	sNew.Steps++
	sNew.Parent = &this
	/*sNew.LastMove = Move{
		Disk:      disk,
		Direction: 1,
	}*/
	temp := sNew.Board[disk][size - 1]
	for i := size - 1; i > 0; i-- {
		sNew.Board[disk][i] = sNew.Board[disk][i - 1]
	}
	sNew.Board[disk][0] = temp
	zeroPos := sNew.Zero
	if zeroPos / size == disk {
		if zeroPos % size == size {
			sNew.Zero = zeroPos + size - 1
		} else {
			sNew.Zero++
		}
	}
	return &sNew
}

func (this State) MoveUp() *State {
	sNew := this
	sNew.Steps++
	sNew.Parent = &this
	/*sNew.LastMove = Move{
		Disk:      -1,
		Direction: 2,
	}*/
	i := sNew.Zero / size
	j := sNew.Zero % size
	Swap(&sNew.Board[i][j], &sNew.Board[i-1][j])
	sNew.Zero -= size
	return &sNew
}

func (this State) MoveDown() *State {
	sNew := this
	sNew.Steps++
	sNew.Parent = &this
	/*sNew.LastMove = Move{
		Disk:      -1,
		Direction: 3,
	}*/
	i := sNew.Zero / size
	j := sNew.Zero % size
	Swap(&sNew.Board[i][j], &sNew.Board[i+1][j])
	sNew.Zero += size
	return &sNew
}

func (this State) GetMoves() *list.List {
	l := list.List{}
	for i := 0; i < size; i++ {
		l.PushBack(this.MoveLeft(i))
		l.PushBack(this.MoveRight(i))
	}
	/*if this.Zero / size > 0 {
		l.PushBack(this.MoveUp())
	}
	if this.Zero / size < 4 {
		l.PushBack(this.MoveDown())
	}*/
	return &l
}

func (this State) Print() {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%4c", toc(this.Board[i][j]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func (this State) IsEqual(s State) bool {
	return this.Board == s.Board && this.Zero == s.Zero
}