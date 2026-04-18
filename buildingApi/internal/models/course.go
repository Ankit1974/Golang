package models

import "strings"

// Domain models / Entities
// Represents real business data

/*
   Used by:

    . Handlers
    . Repositories
    . Services (if added later)
*/

// so basically it's aboute validation and feild matching

/*
     Exactly! You've hit the nail on the head. In the Model layer, you are doing two main things:

      1. Field Matching (The "Bridge")
         The JSON tags (`json:"name"`) act as a bridge.

        Incoming: When a user sends a POST request with {"name": "Go Pro"}, Go uses these tags to know that "name" should go into the Name field of your struct.
        Outgoing: When you send data back, Go uses these tags to name the keys in the JSON response.

      2. Validation (The "Gatekeeper") - The Validate() method is the gatekeeper. It ensures that no "garbage" data enters
	     your system.

         . If a user tries to create a course with a price of -100, your model catches it before it ever reaches your database or repository.

*/

// Course represents a course in the system
type Course struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  int     `json:"price"`
	Author *Author `json:"author"`
}

// Author represents the author of a course
type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// Validate checks if the course has valid data
// Returns true if course is valid, false otherwise
func (c *Course) Validate() (bool, string) {
	if strings.TrimSpace(c.Name) == "" {
		return false, "course name is required"
	}

	if c.Price < 0 {
		return false, "course price must be non-negative"
	}

	if c.Author == nil {
		return false, "course author is required"
	}

	if strings.TrimSpace(c.Author.FullName) == "" {
		return false, "author name is required"
	}

	return true, ""
}

// IsEmpty checks if the course contains meaningful data
func (c *Course) IsEmpty() bool {
	return strings.TrimSpace(c.Name) == ""
}
