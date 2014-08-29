package main

// prepare library:
// $ go get -v -d code.google.com/p/go-tour/tree
import "code.google.com/p/go-tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the cannel ch.
func Walk(t *tree.Tree, ch chan int) {
  _Walk(t, ch)
  close(ch)
}

func _Walk(t *tree.Tree, ch chan int) {
  if t == nil {
    return
  }
  _Walk(t.Left, ch)
  ch <- t.Value
  _Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 cotain the same values.
func Same(t1, t2 *tree.Tree) bool {
  ch1, ch2 := make(chan int), make(chan int)
  go Walk(t1, ch1)
  go Walk(t2, ch2)

  for {
    x, ok1 := <-ch1
    y, ok2 := <-ch2

    if x != y || ok1 != ok2 {
      return false
    }
    if (ok1 && ok2) == false {
      return true
    }
  }
}


func main () {
  fmt.Println( Same(tree.New(1), tree.New(1)) )
  //=> true
  fmt.Println( Same(tree.New(1), tree.New(2)) )
  //=> false
}
