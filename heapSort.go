package main
import (
"fmt"
"time"
"log"
"math/rand"
"github.com/trmchale1/go_py_algos/heap"
)

const Size = 10000000

// builds the list
func buildList() []int {
    fmt.Println("Hello user, sorting 1000 ints...")
    var randList []int
    for i := 0; i < Size; i++{
        randList = append(randList,rand.Intn(Size))
    }
    return randList
}

// prints the list
func printList(list []int){
    l := len(list)
    for i := 0; i < l; i++{
        fmt.Println(list[i])
    }
}

// Build Heap Sort

type minheap struct {
    *heap.Heap
}

// New : returns a new instance of a Heap
func New(input []int) *MinHeap {
	h := &MinHeap{
		&heap.Heap{
			Items: input,
		},
	}

	if len(h.Items) > 0 {
		h.buildMinHeap()
	}

	return h
}

// ExtractMin : removes top element of the heap and returns it
func (h *MinHeap) ExtractMin() int {
	if len(h.Items) == 0 {
		log.Fatal("No items in the heap")
	}
	minItem := h.Items[0]
	lastIndex := len(h.Items) - 1
	h.Items[0] = h.Items[lastIndex]

	// shrinking slice
	h.Items = h.Items[:len(h.Items)-1]

	h.minHeapifyDown(0)

	return minItem
}

func (h *MinHeap) Insert(item int) *minHeap{
    h.tems = append
}

func main() {
    list := buildList()

}
