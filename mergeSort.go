package main
import (
"fmt"
"time"
"log"
"math/rand"
)

const Size = 10000

// builds the list
func buildList() [Size]int {
    fmt.Println("Hello user, sorting 1000 ints...")
    var randList [Size]int
    for i := 0; i < Size; i++{
        randList[i] = rand.Intn(Size)
    }
    return randList
}

// prints the list
func printList(list [Size]int){
    l := len(list)
    for i := 0; i < l; i++{
        fmt.Println(list[i])
    }
}


func main() {
    var w_list[Size]int
    start := time.Now()
    list := buildList()
    // Write mergeSort function call here
    elapsed := time.Since(start)
    log.Printf("Merge Sort took %s", elapsed)
}
