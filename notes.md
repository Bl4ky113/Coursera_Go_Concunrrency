# Concurrency in Go
# University of California, Irvine

Start: 10/01/2024
End:

Sessions:
- 10/01/2024
- 10/03/2024

## What and Why Concurrency

Concurrency is the technique to execute multiple things at the same time, 
it divides itself between parallelism and normal concurrency.

The main difference is that parallelism is mainly made by to different cores on a processor,
meanwhile the concurrency is a way that a single processor can handle multiple actions until completed 
by interrupting and restarting the processes.

Without Concurrency we can only speed up the process by using faster processors with a greater clock rate.
But for this, we also need to speed up every single thing, mainly memory and storage. And that in the long run 
may be a very difficult thing to do. Specialy when we reach processors whom transistors are nearly 5nm.

### Moore's Law
Predicted that every two years, the processor transistor density would double. Making them faster.
This is not a law law, but mearly an observation.

There's a main problem that comes when trying to upgrade and make processors with smaller and smaller transistors, 
making them super dense. And it's called temperature.

The more transistros, more energy they require, and hence by the 3rd law of termodinamic, they generate energy waste called heat.
And independingly from out heat removal system, liquid, fans, or whatever, we can only remove so much heat. This mainly means 
that the processor MUST slow down a little in order to stay in a safe range of temperature, otherwise it would melt and stop working.

## Concepts for Concurrency

A bunch, if not all, of the concepts used in concurrency are from Operating Systems, 
since they MUST have to process all of their processes concurrently, if not at the same time,
define how much and how does each of them can access the resources given to them, mainly due to 
stuff like vistual memory and other stuff.

### Scheduling Algorithmn
The main thing that concurrency needs, not that much if you only use parallelism, 
for example some of this give each task a even amount of time for processing, mainly used in OS.
|
When we switch from one process to other we call that step a Context Switch, where we store the 
context of each process and use them when the process is being executed. 
This switch can sometimes take some loooong time for the speed of processing that we need.

### Threads
Threads are somewhat similar to concurrency in the part htat they need an algorithmn to see which 
thread will execute what and which resources, data, and stuff will it use. 
The man thing is that we can combine both techniques, like multiple processes in a thread, or 
a process which takes multiple threads.
Both stuffs needs to have some sort of safety in order to execute them without problems in 
sharing memory, resources and all the things that could go wrong while doing two things at the same time.

### Bad Interleavings
This is a problem that ocurrs when we are debugging code and we have some problems in 
concurrent code. It happens due to that concurrency is not deterministic and it can change 
in each execution. Hence, when searching for an error, the code execution can alter 
it's behaviour and brake everything, and in other different execution everything can be fine.

### Race Conditions
It's somewhat equal to the previous error, it happends when an output or outcome changes 
when we execute the concurrent code. It mainly happends due to bad or non-existing comunitacion between the processes.
Stuff that we can solve and avoid in Golang with Goroutines.

## Goroutines

Is the implementation of all the mentioned above in Golang. 
Mainly algorithmns that handle the whole processes and their stuff problem.
Whenever we run a go program, Golang will create a main goroutine and execute 
'everything' there, or at least the thing that isn't defined and executed on other goroutine.
We got to keep in mind that if this main goroutine stops, the whole execution of the program stops,
interrupting the other goroutines, which can also be a BIG BIG BIG problem if the stuff that was executed
has problems with beeing stoped, like stoping the writing to a file.

This is what we are going to use in order to make concurrent code. 
The syntaxt is quite simple, we just use the go keyword and call a function:

go function()

Expanding on the exit of the goroutines. When the main routine ends, and one or more extra routines haven't 
is what we call an early exit. We can try to fix this with timing the main exit, sometimes tried with real time timeouts.
But this practice is just bad, and can generate way more problems than before, since the execution is not deterministic.
From there we can see the multiple problem that this generates.

## Sync in Goroutines

We need to have some sort of Syncronization in the routines in order to ensure the correct processing 
of the tasks. The Sync is mainly made by using global events, which might harm a little bit of 
the goal of concurrency, since we have to wait until something is done in order to continue to process tasks.

We can achieve this by using wait groups. Where we can group a bunch of tasks, and use 3 simple methods.
Add() for adding a new task to the group, Done() for anuncing that a task has been completed mainly 
an id or something to know which was completed, and Wait() which will wait unitl all tasks are complete.

This are available in the sync package on the Go stdlib. There's a bunch of other stuff there too, so 
it would be nice to look and see what does that package has.

## Communication between Goroutines

this is a big part of the concurrency, since we need to keep an state, status, information and such shared between the 
tasks in execution. Otherwise, the program tasks could just be each one a different program.

We can achieve this by using Chanels. These are made with make(), and we have to declare the type of data that 
will be used in the channel. Also, in order to use it we will have to use the reciever operator: <- 
Which we will use in order to send data to the channel and fetch data from it as well:

fooChan := make(chan int)

fooChan <- 1 // Send and set data
var bar int = <- fooChan // Fetch and receive data

Since the channel communication and usage needs a sort of syncronization, 
when we fetch the value from one on a routine, we will discreetly do a Wait() for the other channel.
But it wont stop waiting until the routine sends a value by the channel and the waiting routine can use it.

### Buffered Channels
We can define a buffer size for our channel, so we can change the functionality of them.
First, the default buffer for a channel is 0. This is since the routines are blocked or have to wait until some conditions.
In the Sender part, if the buffer is full, the routine will be waiting, 
and in the reciever part, if the buffer is empty, this will be waiting as well.

The size, is defined in the make() capacity arg:
fooChan := make(chan int, 2048)

