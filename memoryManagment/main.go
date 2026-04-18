package main

/*
    Notes :-

	Their are two keywords used in Golang for memory management
	  1. new()
	  2. make()

	  1. new() -
	      . Allocate memory but not insalized
		  . Returns a pointer to it ( memory address )
		  . Allocates zeroed memory
		  . Works with any type

	  2. make() -
	      .	Allocates memory and initializes also
		  . Returns the value, not a pointer ( memory address )
		  .  Works ONLY with:
                # slice
                # map
                # channel
		  . Allocates non - zeroed memory

*/

// Garbage Collector ( GC ) - work Automatically

/*
    Notes : -

	# Stack
      . Fast
      . Function-scoped
      . Automatically cleaned up
      . No GC overhead

	# Heap
      . Slower than stack
      . Used when data outlives the function
      . Managed by GC


	  Go compiler decides:
            “Does this variable need to live beyond this function?”

       If yes → heap
       If no → stack


*/
