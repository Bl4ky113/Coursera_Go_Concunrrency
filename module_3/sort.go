package main

import (
    "fmt"
    "math"
    "sort"
)

func main () {
    var startIndex, endIndex, numArrLength, sliceSize, numSlices int

    numArr := make([]int, 0, 0)
    resultArr := make([]int, 0, 0)
    sortedArrChan := make(chan *[]int, 4)

    numArrLength = getUsrInput(&numArr)
    sliceSize = int(math.Ceil(float64(numArrLength) / 4.0))

    for endIndex < numArrLength || numSlices < 4 {
        startIndex = endIndex

        if sliceSize > (numArrLength - endIndex) {
            endIndex += (numArrLength - endIndex)
        } else {
            endIndex += sliceSize
        }

        
        arraySlice := numArr[startIndex:endIndex]
        fmt.Println(arraySlice)
        go sortArrayRoutine(&arraySlice, sortedArrChan)

        numSlices++
    }

    for numSlices > 0 {
        sortedArr := <- sortedArrChan
        resultArr = append(resultArr, *sortedArr...)
        numSlices--
    }        

    sortArray(&resultArr)
    fmt.Println(resultArr)
    
    return
}

func getUsrInput (inputArr *[]int) int {
    var arrLength int

    fmt.Printf("Enter the size of the list to sort:\t")
    fmt.Scanf("%d", &arrLength)

    for i := 0; i < arrLength; i++ {
        var usrInput int

        fmt.Printf("Enter number N.%d:\t", i + 1)
        fmt.Scanf("%d", &usrInput)

        (*inputArr) = append((*inputArr), usrInput)       
    }

    return arrLength
}

func sortArray (numArr *[]int) {
    // Let's be real, the last 'real' thing to check in this exercise 
    // is how does the program sorts the slice
    sort.Slice((*numArr), func(i, j int) bool { return (*numArr)[i] < (*numArr)[j] })
    return
}

func sortArrayRoutine (numArr *[]int, outputChan chan *[]int) {
    sortArray(numArr)

    outputChan <- numArr
}
