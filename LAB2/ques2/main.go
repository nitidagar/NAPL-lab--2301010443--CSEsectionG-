//package main

//import "fmt"

/*func main() {

	// Create a slice
	numbers := []int{10, 20, 30}

	fmt.Println("Initial Slice:", numbers)

	// 1. Add / Append
	var num int

	fmt.Print("Enter a number to add: ")
	fmt.Scan(&num)

	numbers = append(numbers, num)

	fmt.Println("After Add:", numbers)

	// 2. Remove by Index
	var index int

	fmt.Print("Enter index to remove: ")
	fmt.Scan(&index)

	if index >= 0 && index < len(numbers) {
		numbers = append(numbers[:index], numbers[index+1:]...)
		fmt.Println("After Remove:", numbers)
	} else {
		fmt.Println("Invalid index")
	}

	// 3. Update
	var updateIndex int
	var newValue int

	fmt.Print("Enter index to update: ")
	fmt.Scan(&updateIndex)

	if updateIndex >= 0 && updateIndex < len(numbers) {
		fmt.Print("Enter new value: ")
		fmt.Scan(&newValue)

		numbers[updateIndex] = newValue

		fmt.Println("After Update:", numbers)
	} else {
		fmt.Println("Invalid index")
	}

	//(2) uses a map to store key–value pairs (e.g., subject → marks) and performs insert, delete, and lookup operations, printing the map after each operation.

	// Create a map
	marks := map[string]int{
		"Maths":   85,
		"Science": 90,
		"English": 80,
	}

	fmt.Println("Initial Map:", marks)

	// 1. Insert
	var subject string
	var mark int

	fmt.Print("Enter subject to insert: ")
	fmt.Scan(&subject)

	fmt.Print("Enter marks: ")
	fmt.Scan(&mark)

	marks[subject] = mark

	fmt.Println("After Insert:", marks)

	// 2. Delete
	var deleteSubject string

	fmt.Print("Enter subject to delete: ")
	fmt.Scan(&deleteSubject)

	delete(marks, deleteSubject)

	fmt.Println("After Delete:", marks)

	// 3. Lookup
	var lookupSubject string

	fmt.Print("Enter subject to lookup: ")
	fmt.Scan(&lookupSubject)

	value, found := marks[lookupSubject]

	if found {
		fmt.Println("Marks of", lookupSubject, ":", value)
	} else {
		fmt.Println("Subject not found")
	}

	fmt.Println("After Lookup:", marks)
}
*/