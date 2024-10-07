# Concurrency in Go
# University of California, Irvine

Start: 10/01/2024
End:

Sessions:
- 10/01/2024
- 10/03/2024
- 10/07/2024

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

We can iter through a channel with the range keyword in a for loop. Allowing us to do stuff with each thing that 
we receive from that channel. This loop might be infinite, that's why we can close the channel by using the 
funciton close, passing the channel as an argument. It's not that important to close the channel, like a file, 
but it is if you are using a for loop for fetching the channel data.

### Deadlock
Is a problem where we have routines that have a mutual or circular dependency. 
Dependency meaning that one should finish in order to the other to start. 
Meaning that these will stay forever waiting for the other to finish. 
Go will help us a little, trowing an fatal error if all routines, including the main one are blocked. But only if ALL routines are blocked.
This is a very hard stuff to detect if it happens, so be careful whenever you implement a locked concurrent algorithmn.

### Buffered Channels
We can define a buffer size for our channel, so we can change the functionality of them.
First, the default buffer for a channel is 0. This is since the routines are blocked or have to wait until some conditions.
In the Sender part, if the buffer is full, the routine will be waiting, 
and in the reciever part, if the buffer is empty, this will be waiting as well.

The size, is defined in the make() capacity arg:
fooChan := make(chan int, 2048)

### Select statement
When we have multiple channels that we have to get, we can wait for both and use them. 
But if we need only one channel to do something while we are waiting for 2 or more we can use 
the Select Statement:

select {
    case foo <- c1:
        ...
    case foo <- c2:
        ...
    case c3 <- bar:
        ...
    ...
    default:
        ...
}

Keep in mind that we can also, while waiting for data, have a case for sending data to another channel,
which might work as an default case if the channel isn't closed.
But if we need a normal default case, we can also define it as well.

A good use of the select statement is the usage of abort channels, whenever we are waiting for something and 
we get an interrupt or anything signal we han handle safely the exit of the function and the concurrency.

### Mutual Exclusion - MutEx

When two routines share variables, there might be errors when one of then changes the variable and interferes the 
execution of the other routine. This means that the functions are not concurrent-safe with some variable or data,
for example most of DB writes, updates and deletes are not concurrent-safe. This mainly happens with stuff 
that firstly reads and then writes data in the machine level of code.

A simple fix for this is just not using functions that write a variable at the same time in concurrency, 
or restrict this possible interleavings.

We can do this restrictions with MutEx, available in the sync package, which blocks the execution on routines. 
This is made by using the method Lock() where a routine will execute code that will 
change the value of a shared variable, blocking ALL other routines which has the mutex instance
with a call to Lock(). Then after the Locking routine is over, calls Unlock() to let another 
routine Lock the variable and repeat the cycle until every routine ends.

## Sync Once

When we have a group of routines to execute, we might want to have them share some stuff like variables and 
what not. This things are generally defined in a init enviroment like the main function, but maybe when we have 
to execute them from another package we might need to use another function. The problem with this function is that 
it might not sync well with the routes, so we can use sync.Once and it's method Do().
Which allow us to block routines until a function is done executing, generally an init function.

## Dinning Phisolophers

The hypotesis of the problem is that we have 5 Phisolophers in a circular table, each one has a bowl of rice, 
also everyone has 1 chopstick placed at their left side. But in order to eat the rice, they need 
to use 2 chopstick. What they can do is that they can grab the chopstick at their right, the chopstick of their neighbor.
Blocking them to eat the rice.

Now the big problem is that if everyone picks their chopstick,
there wouldn't be any chopstick left to complete the pair. So everyone will starve.
Making a Deadlock.

One solution to the problem is that each Phisolophers would take the chopstick with the lowest index number available to them.
So 0th would take the right chopstick of 4th and 3rd would take 4th's left chopstick. Meaning that 4th would starve.
This solution doesn't deadlock itself, but will starve the 4 routine.

One that I thought was that we cycle in pairs in a group way, using module of the number of Phisolophers.
