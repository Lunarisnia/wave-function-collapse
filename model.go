package wavefunctioncollapse

import (
	"fmt"
	"math/rand"
)

// For each cell, create a boolean array representing the domain of that variable. The domain has one entry per tile, and they are all initialized to true.
// A tile is in the domain if that entry is true.

// While there are still cells that have multiple items in the domain:

// Pick a random cell using the “least entropy” heuristic.

// Pick a random tile from that cell’s domain, and remove all other tiles from the domain.

// Update the domains of other cells based on this new information, i.e. propagate cells. This needs to be done repeatedly as changes to those cells may have further implications.

type Domain string

const (
	A Domain = "A"
	B Domain = "B"
)

type Model struct {
	Output []Domain
	Wave   [][]Domain
}

func (m *Model) CollapseRandomly() int {
	randomIndex := rand.Intn(len(m.Wave))
	if !m.Collapsed(randomIndex) {
		collapsedIndex := rand.Intn(2)
		m.Wave[randomIndex] = m.Wave[randomIndex][collapsedIndex : collapsedIndex+1]
		return randomIndex
	}

	return -1
}

func (m *Model) Collapsed(i int) bool {
	return len(m.Wave[i]) == 1
}

func (m *Model) ConstraintPropagating(i int) {
	// remove a in the right if its collapsed to A
	if m.Collapsed(i) && m.Wave[i][0] == A {
		if i+1 < len(m.Wave) {
			fmt.Println("Valid right")
			for d, domain := range m.Wave[i+1] {
				if domain == B {
					fmt.Println(m.Wave[i+1][d:d+1], "=====")
					m.Wave[i+1] = m.Wave[i+1][d : d+1]
				}
			}
		}
	}
}

func NewModel() *Model {
	model := &Model{
		Wave: make([][]Domain, 5),
	}

	for i := range model.Wave {
		model.Wave[i] = []Domain{A, B}
	}
	// 1. Declare array of length 5x2
	// 2. Fill each rows with 2 Possible values(domain)
	// 3. Collapsing means selecting one of two possible values
	// 4. After a random value is collapsed we need to do constraint checking
	// A cannot be followed by another A

	return model
}

func Run() {
	model := NewModel()
	for {
		fullyCollapsed := true
		for i := range model.Wave {
			if !model.Collapsed(i) {
				fmt.Println("I: ", model.Wave)
				fullyCollapsed = false
				break
			}
		}
		if fullyCollapsed {
			fmt.Println("Fully Collapsed")
			fmt.Println(model.Wave)
			return
		}
		collapsedIndex := model.CollapseRandomly()
		if collapsedIndex != -1 {
			model.ConstraintPropagating(collapsedIndex)
		}

	}

	arr := []int{1, 2}
	fmt.Println(arr[0:1], arr[1:2])
}
