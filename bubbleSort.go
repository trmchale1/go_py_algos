package main
import (
"fmt"
"math/rand"
"time"
"log"
)

// size 10k 161ms
// size 100k 20s
const Size = 100000

// builds an array of size Size and returns it

func buildList() [Size]int {
    fmt.Println("Hello user, sorting 1000 ints...")
    var randList [Size]int
    for i := 0; i < Size; i++{
        randList[i] = rand.Intn(Size)
    }
    return randList
}

// prints our list
func printList(list[Size] int){
    l := len(list)
    for i := 0; i < l; i++{
        fmt.Println(list[i])
    }
}

// bubbleSort O(n^2)
func bubbleSort(list[Size] int) [Size] int {
    l := len(list)
    for i := 0; i < (l-1); i++ {
        for j := 0; j < l-1-i; j++ {
            if(list[j+1] < list[j]){
                temp := list[j]
                list[j] = list[j+1]
                list[j+1] = temp
            }
        }
    }
    return list
}

func main(){
    start := time.Now()
    var list[Size]int
    list = buildList()
    list = bubbleSort(list)
    printList(list)
    elapsed := time.Since(start)
    log.Printf("Merge Sort took %s", elapsed)
}
