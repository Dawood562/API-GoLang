package structs

// Determine Data Structure
type Xyz struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Number      int    `json:"number"`
	SomeBoolean bool   `json:"someBoolean"`
}

// Structure specifically for PATCH requests to ensure I can tell which data needs to be patched
type XyzPatch struct {
	Id          *string
	Title       *string
	Description *string
	Number      *int
	SomeBoolean *bool
}
