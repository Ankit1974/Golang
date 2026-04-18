package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
)

// %w is used to wrap an error so it can be inspected later.
/*
   1. m%w is used inside fmt.Errorf to wrap an existing error while adding more context.
   code - fmt.Errorf("something failed: %w", err)

*/

/*
     Notes : -

	 # Why Go even needs %w
        In production code you want two things at the same time:
             1. Human-readable context
             2. Ability to programmatically detect the original error
          Without %w, you lose #2.
*/

/*
		         Scenario

                . Frontend developer calls your API
                . Something goes wrong on the backend
                . You handle it correctly (dev-friendly + user-safe)

              Example: Frontend calls GET /profile
                  . What can go wrong?
                  . DB file missing
                  . Config missing
                  . Internal error

            We’ll use fs.ErrNotExist as the root cause.

                    Backend code (production-style)

                    Step 1: Low-level (real failure).
*/

func loadUserProfile(userID string) error {
	_, err := os.ReadFile("users.db")
	if err != nil {
		return fmt.Errorf("load user profile failed: %w", err)
	}
	return nil
}

/*  Step 2: HTTP handler (decide response) */

func getProfileHandler(w http.ResponseWriter, r *http.Request) {
	err := loadUserProfile("123")
	if err != nil {

		// Case 1: Known error → client-safe response
		if errors.Is(err, fs.ErrNotExist) {
			http.Error(
				w,
				"Profile not found",
				http.StatusNotFound,
			)
			return
		}

		// Case 2: Unknown error → generic message
		log.Println("internal error:", err)

		http.Error(
			w,
			"Something went wrong. Please try again later.",
			http.StatusInternalServerError,
		)
		return
	}

	w.Write([]byte("Profile data"))
}

/*
    What frontend developer sees 👀

           1. HTTP response
              . Status: 404
              . Body: "Profile not found"

              Frontend developer understands:
                 1. API worked
                 2. Resource not found
                3. Not a crash

    What backend developer sees 👨‍💻
              1. Logs
              2. load user profile failed: open users.db: no such file or directory

  Full context + root cause preserved 🎯
*/

/*
   Notes : -

  1. loadUserProfile → backend logic
  2. getProfileHandler → backend API
*/
