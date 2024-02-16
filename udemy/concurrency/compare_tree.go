package main

import (
"fmt"
"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	
	if t == nil {
		return
	}
	
	var next []*tree.Tree
	next = append(next,t)
	
	for len(next) > 0 {
		p := next[len(next)-1]

		ch <- p.Value
	
		if p.Left != nil {
			next = append(next,p.Left)
		} 
		
		if p.Left != nil {
			next = append(next,p.Right)
		}
		
		next = next[:len(next)-1]
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1,ch1)
	go Walk(t2,ch2)
	
	for {
		a := <-ch1
		b := <-ch2
		
		if a != b {
			fmt.Println(a)
			fmt.Println(b)
			fmt.Println("")
			return false 
		}
	}

	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
    ch1 := make(chan int)
	ch2 := make(chan int)
	
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	fmt.Println("T1:")
	for c := range <- ch1 {
		fmt.Printf("%d ", c)
	}
	fmt.Println("")
	
	fmt.Println("T2:")
	for c := range <- ch2 {
		fmt.Printf("%d ", c)
	}
	fmt.Println("")
	
	fmt.Println(Same(t1, t2))

}
