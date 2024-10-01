# Concurrency in Go
# University of California, Irvine

Start: 10/01/2024
End:

Sessions:
- 10/01/2024

## What and Why Concurrency

Concurrency is the techinique to execute multiple things at the same time, 
it divides itself between parallelism and normal concurrency.

The main difference is that parallelism is mainly made by to different cores on a processor,
meanwhile the concunrrency is a way that a single processor can handle multiple actions until completed 
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

