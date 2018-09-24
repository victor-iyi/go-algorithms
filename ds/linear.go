package ds

// LinearDS interface contains common method signatures
// for implemented data structures.
type LinearDS interface {
	Data
	Push(data Data)      // Adding elements to a data structure.
	Peek() (Data, error) // Returns the first element of a data structure.
	Pop() (Data, error)  // Removes element based on data structure principle.
	Count() int          // Count number of elements is in a data structure.
	Empty() bool         // Checks if a data structure is empty or not.
}

// Data contains element types in data structure.
type Data interface {
}
