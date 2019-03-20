package main

import (
	"fmt"
)

//main function
func main() {
	//Declare an int variable
	var x int = 10
	//change value of x using pointer
	changeValueUsingPointer(&x)
	fmt.Println(x)

	//Create People object
	var people People

	//Store returned value List
	peopleInstance,
		stringValue,
		intValue,
		unsignedEightBitIntegerValue,
		unsignedSixteenBitIntegerValue,
		unsignedThirtyTwoBitIntegerValue,
		unsignedSixtyFourBitIntegerValue,
		signedEightBitIntegerValue,
		signedSixteenBitIntegerValue,
		signedThirtyTwoBitIntegerValue,
		signedSixtyFourBitIntegerValue,
		floatThirtyTwoValue,
		floatSixtyFourValue,
		boolValue := returnTwoValue(people)
	fmt.Println(
		"get value by attribute: ", peopleInstance.name, "\n",
		"String: ", stringValue, "\n",
		"Integer: ", intValue, "\n",
		"Unsigned 8-bit integers (0 to 255): ", unsignedEightBitIntegerValue, "\n",
		"Unsigned 16-bit integers (0 to 65535): ", unsignedSixteenBitIntegerValue, "\n",
		"Unsigned 32-bit integers (0 to 4294967295): ", unsignedThirtyTwoBitIntegerValue, "\n",
		"Unsigned 64-bit integers (0 to 18446744073709551615)): ", unsignedSixtyFourBitIntegerValue, "\n",
		"Signed 8-bit integers (-128 to 127): ", signedEightBitIntegerValue, "\n",
		"Signed 16-bit integers (-32768 to 32767): ", signedSixteenBitIntegerValue, "\n",
		"Signed 32-bit integers (-2147483648 to 2147483647): ", signedThirtyTwoBitIntegerValue, "\n",
		"Signed 64-bit integers (-9223372036854775808 to 9223372036854775807): ", signedSixtyFourBitIntegerValue, "\n",
		"Float32: ", floatThirtyTwoValue, "\n",
		"Float64: ", floatSixtyFourValue, "\n",
		boolValue)

	//Declare a constant
	const HEIGHT int = 5
	fmt.Println("Constant value:", HEIGHT)

	str := ifElse(26)
	fmt.Println(str)

	//print values from an array using for loop

	forLoop()
	//print values from an array using for while loop
	whileLoopByCustomizingForLoop()
	//Array Operations
	arrayOperations()
	//Map Operations
	mapExample()
	//Calling a recursive method
	recursionExample(1)

	//Craeting an instance of SinglyLinkedListNode struct
	var head SinglyLinkedListNode
	//Initializing the
	head.data = 1
	//Add node
	insert(&head, 2)
	insert(&head, 3)
	insert(&head, 4)
	//print a node data
	fmt.Println("print a node data", head.next.next.data)

	//Printing Node data of a Singly LinkedList
	fmt.Println("Printing node data of Singly Linked List Recursively:", head.data)
	printDataOfSinglyLinkedListUsingRecursion(&head)

}

// usages of pointer; manipulate value of a variable using pointer.
func changeValueUsingPointer(x *int) {
	*x = 12
}

//function that returns n values
func returnTwoValue(people People) (People, string, int, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64, bool) {
	people.name = "Shahidul Islam"
	people.email = "www.zeromsi2@gmail.com"
	people.phone = 1743756128
	return people, ";Man in form", 10, 255, 65535, 4294967295, 18446744073709551615, 127, 32767, 2147483647, 9223372036854775807, 12.343576, 12.353453465356656, true
}

//Usage of condition
func ifElse(x int) string {
	if x > 10 && x < 20 {
		return "Greater than 10 but smaller that 20"
	} else if (x > 20 && x < 30) && x == 25 {
		return "Greater than 20 but smaller that 30 and the value is 25"
	} else if x > 20 && x < 30 {
		if x == 26 {
			return "Value is 26"
		}
		return "Greater than 20 but smaller that 30"
	} else {
		return "None of this"
	}

}

//For Loop
func forLoop() {
	values := [9]int{1, 2, 3, 5, 6, 7, 8, 9, 10}
	fmt.Println("Let's Run a For Loop:")
	for i := 0; i < len(values); i++ {
		fmt.Printf("Array element at %d = %d\n", i, values[i])
	}
}

//While Loop
func whileLoopByCustomizingForLoop() {
	var i int
	values := [9]int{1, 2, 3, 5, 6, 7, 8, 9, 10}
	fmt.Println("Let's Run a While Loop:")
	for i < len(values) {
		fmt.Printf("Array element at %d = %d\n", i, values[i])
		i++
	}
}

//Array operations

func arrayOperations() {
	values := []int{1, 2, 3, 5, 6, 7, 8, 9, 10}
	fmt.Println("Let's Print All values from an array")
	fmt.Println(values)
	fmt.Println("Let's Add an element:")
	values = append(values, 11)
	fmt.Println(values)
}

//Map operation
func mapExample() {
	fmt.Println("Print student(key=3)name from map:")
	student := make(map[int]string)
	student[1] = "Student One"
	student[2] = "Student Two"
	student[3] = "Student Three"
	student[4] = "Student Four"
	//print student name from map
	fmt.Println(student[3])
	fmt.Println("Delete Student(key=3) from map:")
	delete(student, 3)
	fmt.Println("Print all Student:")
	fmt.Println(student)
}

//Recursion Example
func recursionExample(i int) {
	if i < 10 {
		fmt.Println("Recursively  Printing:", i)
		recursionExample(i + 1)
	}
}

//Define a struct;Derived data type
type People struct {
	name  string
	email string
	phone int
}

//Define a SinglyLinkedList Node(struct)
type SinglyLinkedListNode struct {
	data int
	next *SinglyLinkedListNode
}

//insert a node into Singly Linked List
func insert(node *SinglyLinkedListNode, data int) {
	if node.next == nil {
		temp := &SinglyLinkedListNode{
			data: data,
			next: nil,
		}
		node.next = temp
	} else {
		insert(node.next, data)
	}
}
//Printing node data of Singly Linked List Recursively
func printDataOfSinglyLinkedListUsingRecursion(node *SinglyLinkedListNode) {
	if node.next != nil {
		fmt.Println("Printing node data of Singly Linked List Recursively:", node.next.data)
		printDataOfSinglyLinkedListUsingRecursion(node.next)

	}
}

