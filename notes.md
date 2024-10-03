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

There's a main problemn that comes when trying to upgrade and make processors with smaller and smaller transistors, 
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
Both stuffs needs to have some sort of safety in order to execute them without problemns in 
sharing memory, resources and all the things that could go wrong while doing two things at the same time.

### Interleavings
This is a problemn that ocurrs when we are debugging code and we have some problemns in 
concurrent code. It happens due to that concurrency is not deterministic and it can change 
in each execution. Hence, when searching for an error, the code execution can alter 
it's behaviour and brake everything, and in other different execution everything can be fine.

### Race Conditions
It's somewhat equal to the previous error, it happends when an output or outcome changes 
when we execute the concurrent code. It mainly happends due to bad or non-existing comunitacion between the processes.
Stuff that we can solve and avoid in Golang with Goroutines.

## Goroutines

Is the implementation of all the mentioned above in Golang. 
Mainly algorithmns that handle the whole processes and their stuff problemn. 

This is what we are going to use in order to make concurrent code.
