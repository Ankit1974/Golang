package racecondtion

/*
  Simple Definition:
      A race condition occurs when two or more goroutines access shared data at the same time, and at least one of them modifies it.

	  The final result then depends on:
      👉 which goroutine runs first

      Super Simple Meaning
      Think of it like:

      “Two goroutines racing to change the same value.”
      Whoever wins the race changes the result.

         Real-World Analogy

           Imagine:
            Two people using the same bank account
            Both try to withdraw money at the same time
            Account balance = 1000
            Person A withdraws 300
            Person B withdraws 400

*/

// go run --race ( this is used to find the race condition in the code)
