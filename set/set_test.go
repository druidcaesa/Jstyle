package set

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	s := New[int]()
	s.Put(2)
	fmt.Println(s)
}

func TestForEach(t *testing.T) {
	s := New[int]()
	s.Put(2)
	s.Put(3)
	s.Put(5)
	s.Put(6)
	s.Put(10)
	s.Put(2)
	s.Put(32)
	s.Put(1)
	s.ForEach(func(k int) {
		fmt.Println(k)
	})
}
