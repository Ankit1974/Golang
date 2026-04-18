package pipeline
'
'/*

  🔹 What is a Pipeline in Go?

      Simple Definition
	  
	  A Pipeline is a chain of stages where data flows through multiple goroutines using channels.

      Each stage:

      1. Takes input from a channel

      2. Processes it

      3. Sends output to next channel

      Visual Idea

         Input → Stage1 → Stage2 → Stage3 → Output

      Data moves like water through pipes

      👉 Hence the name Pipeline

*/

/*

   Basic Structure of Pipeline :- 

  1. Every pipeline stage has:
  2. Input channel
  3. Processing logic
  4. Output channel

*/


 // STAGE 1 :- Generator

 func generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}


// STAGE 2 :- Square

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}


// STAGE 3 :- Printer

func printer(in <-chan int) {
	for n := range in {
		fmt.Println(n)
	}
}


// MAIN FUNCTION

func main() {
	// Pipeline: Generator → Square → Printer
	gen := generator(1, 2, 3, 4, 5)
	sq := square(gen)
	printer(sq)
}


/*
   Key Properties of Pipelines

 1️⃣ Stages are Independent
    Each stage:
    Runs in its own goroutine
    Doesn’t know about others

 2️⃣ Channels Connect Stages
    Channels act as pipes between stages.

 3️⃣ Backpressure
    If next stage is slow:
    Previous stage automatically waits
    No extra code needed