package main

import (
    "fmt"
    "sync"
)

const (
    NUM_CHOPSTICKS = 5
    NUM_PHILOSOPHERS = 5
    MAX_EATING_PHILOSOPHERS = NUM_CHOPSTICKS / 2
)

type Chopstick struct {
    sync.Mutex
}

type Philosopher struct {
    leftChopstick *Chopstick
    rightChopstick *Chopstick
    food int
    number int
    done bool
}

var LunchInProgress bool = true

func (phi *Philosopher) eatRice (outGroup *sync.WaitGroup, outChannel chan bool) (bool) {
    if phi.food <= 0 {
        if !phi.done {
            fmt.Printf("Phi. #%d finished eating fried rice\n", phi.number)
            outChannel <- true
            phi.done = true
        }
        outGroup.Done()
        return true
    } 

    phi.leftChopstick.Lock()
    phi.rightChopstick.Lock()

    phi.food--
    fmt.Printf("Phi. #%d is eating, %d servings left\n", phi.number, phi.food)

    phi.leftChopstick.Unlock()
    phi.rightChopstick.Unlock()

    outChannel <- false
    outGroup.Done()
    return false
}

func main () {
    chopArr := make([]Chopstick, NUM_CHOPSTICKS, NUM_CHOPSTICKS)
    phiArr := make([]Philosopher, NUM_PHILOSOPHERS, NUM_PHILOSOPHERS)

    createChopsticks(&chopArr)
    createPhilosophers(&phiArr, &chopArr)

    mainWg := new(sync.WaitGroup)
    mainWg.Add(1)
    
    go handlePhilosophersLunch(&phiArr, mainWg)

    mainWg.Wait()
    return
}

func createChopsticks (chopArr *[]Chopstick) {
    for i := 0; i < NUM_CHOPSTICKS; i++ {
        (*chopArr)[i] = *new(Chopstick)
    }

    return
}

func createPhilosophers (phiArr *[]Philosopher, chopArr *[]Chopstick) {
    for i := 0; i < NUM_PHILOSOPHERS; i++ {
        (*phiArr)[i] = *new(Philosopher)
        (*phiArr)[i].food = 3
        (*phiArr)[i].number = i + 1
        (*phiArr)[i].done = false
        (*phiArr)[i].leftChopstick = &(*chopArr)[i]
        (*phiArr)[i].rightChopstick = &(*chopArr)[((i - 1) + 5) % NUM_PHILOSOPHERS]
    }

    return
}

func handlePhilosophersLunch (phiArr *[]Philosopher, outerWg *sync.WaitGroup) {
    defer outerWg.Done()

    eatingPhiGroup := new(sync.WaitGroup)
    eatingPhiChannel := make(chan bool, 2)
    exitChannel := make(chan bool)

    numPhiDoneEating := 0

    go handleMealsProgress(&numPhiDoneEating, eatingPhiChannel, exitChannel)

    for LunchInProgress && numPhiDoneEating < NUM_PHILOSOPHERS {
        for i := 0; i < NUM_PHILOSOPHERS; i++ {
            for j := i; (j - i) <= MAX_EATING_PHILOSOPHERS; j += 2 {
                eatingPhiGroup.Add(1)
                go (*phiArr)[j % NUM_PHILOSOPHERS].eatRice(eatingPhiGroup, eatingPhiChannel)
            }
            
            eatingPhiGroup.Wait()

            if !LunchInProgress || numPhiDoneEating >= NUM_PHILOSOPHERS {
                break
            }

            fmt.Printf("Round Done!\t%d\t%d\t%v\n", numPhiDoneEating, NUM_PHILOSOPHERS, LunchInProgress)

            select {
            case <- exitChannel:
                LunchInProgress = false
            default:
            }
        }
    }

    fmt.Printf("\nLunch Done!\n")
    return
}

func handleMealsProgress (numPhiDoneEating *int, eatingPhiChannel chan bool, exitChannel chan bool) {
    for *numPhiDoneEating < NUM_PHILOSOPHERS {
        select {
        case finishedMeal := <- eatingPhiChannel:
            if finishedMeal {
                *(numPhiDoneEating)++
            }
        }
    }
    
    exitChannel <- true
    return
}
