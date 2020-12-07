package main
import (
"fmt"
"time"
"log"
"math/rand"
)


// 10k 25ms
// 100k 500ms
// 1m 5s
// 10m 50s
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

func mergeSort(slice []int) []int {
    if len(slice) < 2 {
        return slice
    }
    m := (len(slice)) / 2
    return merge(mergeSort(slice[:m]), mergeSort(slice[m:]))
}

func merge(left, right []int) []int {
    size, i, j := len(left) + len(right), 0,0
    slice := make([]int, size, size)

    for k := 0; k < size; k++ {
        if i > len(left)-1 && j <= len(right)-1 {
            slice[k] = right[j]
            j++
    }  else if j > len(right)-1 && i <= len(left)-1{ 
        slice[k] = left[i]
        i++
    } else if left[i] < right[j] {
        slice[k] = left[i]
        i++
    } else {
        slice[k] = right[j]
        j++
    }
}
    return slice
}

func main() {
    start := time.Now()
    list := buildList()
    // Write mergeSort function call here
    printList(list)
    list = mergeSort(list)
    printList(list)
    elapsed := time.Since(start)
    log.Printf("Merge Sort took %s", elapsed)
}
