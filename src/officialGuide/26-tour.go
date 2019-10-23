package main

import (
    "golang.org/x/tour/tree"
    "fmt"
)

func Walk (t *tree.Tree, ch chan int) {
    if t.Left != nil {
        Walk(t.Left, ch)
    }
    ch <- t.Value
    if t.Right != nil {
        Walk(t.Right, ch)
    }
}

func TestWalk() {
    ch := make(chan int, 10)
    k := 2
    go Walk(tree.New(k), ch)
    for i := 1; i <= 10; i++ {
        val := <-ch
        if i * k != val {
            fmt.Println("FAIL: Walk() goes wrong.")
            return
        }
    }
    fmt.Println("PASS: Walk() goes well")
}

func Same(t1, t2 *tree.Tree) bool {
    ch1 := make(chan int, 10)
    ch2 := make(chan int, 10)
    go Walk(t1, ch1)
    go Walk(t2, ch2)
    for i := 0; i < 10; i++ {
        n1, n2 := <-ch1, <-ch2
        if n1 != n2 {
            return false
        }
    }
    return true
}

func main(){
    ch := make(chan int, 10)
    k := 1
    go Walk(tree.New(k), ch)
    for i := 0; i < 10; i++ {
        fmt.Println(<-ch)
    }
    TestWalk()

    same := Same(tree.New(1), tree.New(1))
	if same {
		fmt.Println("PASS : they are same")
	} else {
		fmt.Println("FAIL : they should be same.")
	}

	same = Same(tree.New(1), tree.New(2))
	if !same {
		fmt.Println("PASS : they are not same")
	} else {
		fmt.Println("FAIL : they should be not same")
	}
}

