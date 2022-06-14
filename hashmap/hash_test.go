package hashmap

import (
	"fmt"
	"testing"
)

func TestHashMap(t *testing.T) {
	hashMap := NewHashMap[string, string]()
	hashMap.Put("t", "1")
	hashMap.Put("t2", "2`")
	hashMap.Put("t3", "3")
	hashMap.Put("t4", "4")

	hashMap.ForEach(func(key string, val string) {
		fmt.Printf("key--------->%s\n,val--------->%s\n", key, val)
	})

}