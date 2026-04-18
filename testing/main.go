package testing

// we use test.go library to write test cases

// Type of testing ing Golang

/*
  1. Table - driven tests
  2. Sub tests
  3. Benchmark tests
  4. Test coverge
*/

/*
   Use them when the SAME logic must be tested with MULTIPLE inputs.

      ✅ Use table-driven tests if:

          . Same function
          . Many scenarios / edge cases
          . Business rules vary by input
          . Validation logic exists
*/

/*

   Subtests are NOT optional if you use table-driven tests.

     ✅ Use subtests when:

         . You want named test cases
         . You want clear failure output
         . You want isolated test execution

*/

/*
   This is where most juniors get it wrong.

   ✅ Use benchmarks ONLY when:

       . Function runs very frequently
       . Code is CPU or memory heavy
       . Performance affects users
       . You are comparing two approaches
*/
