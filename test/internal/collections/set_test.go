package collections

import (
	"fmt"
	"github.com/danjsg/simpleauth/internal/collections"
	"testing"
)

func TestSet(t *testing.T) {
	set := collections.HashSet[string]()
	set.Add("A")
	set.Add("B")
	set.Add("C")
	set.Add("D")
	fmt.Printf("Set has %d elements\n", set.Size())
	fmt.Println("Set elements: ")
	iterator := set.Iterator()
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
	set.Remove("A")
	set.Remove("B")
	fmt.Printf("Set has %d elements\n", set.Size())
	fmt.Println("Set elements: ")
	iterator = set.Iterator()
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
	fmt.Printf("Set still has %d elemenets\n", set.Size())
	fmt.Printf("Does set contain %s? %s\n", "A", yesNoMapper(set.Contains("A")))
	fmt.Printf("Does set contain %s? %s\n", "C", yesNoMapper(set.Contains("C")))
	set.Add("C")
	set.Add("D")
	fmt.Printf("Set has %d elements\n", set.Size())
	fmt.Println("Set elements: ")
	iterator = set.Iterator()
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}

}

func yesNoMapper(b bool) string {
	if b {
		return "Yes"
	} else {
		return "No"
	}
}
