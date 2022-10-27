package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
 * Oct 25, 2022
 * ptp-glng033
 * GLNG033ONL1006
 * iota
 * simplifies constants predefined value
 * iota value: index of constant
 * ref: https://go.dev/ref/spec#Iota
 */

//-------------------------------------------------------------
/* [START] BASIC CONST() USAGE */
// in const(), every variables which is not inline-assigned
// will be assigned with same value from previous expression:
// const (
/* VAR 		EXPR	RESULT*/
// 	MYCONST1 = 1 //1
// 	MYCONST2     //1
// 	MYCONST3     //1
// 	MYCONST4     //1
// 	MYCONST5     //1
// 	MYCONST6     //1
// )

// const (
/* VAR 		EXPR	RESULT*/
// 	MYCONST1 = 0     //0
// 	MYCONST2         //0
// 	MYCONST3 = 5     //5
// 	MYCONST4         //5
// 	MYCONST5 = 1 + 1 //2
// 	MYCONST6         //2
// )

/* [END] BASIC CONST() USAGE */

//-------------------------------------------------------------
/* [START] IOTA USAGE */
// with basic pricipal of const() assignation process,
// iota could be used to make constant values unique.
// you can name it as enumeration as well (?)
// const (
/* VAR 		EXPR	RESULT*/
// 	MYCONST1 = iota //0
// 	MYCONST2     //1
// 	MYCONST3     //2
// 	MYCONST4     //3
// 	MYCONST5     //4
// 	MYCONST6     //5
// )

// const (
// 	/* VAR 		EXPR	RESULT*/
// 	MYCONST1 = 0        //0
// 	MYCONST2            //0
// 	MYCONST3 = iota     //2
// 	MYCONST4            //3
// 	MYCONST5 = iota * 2 //8
// 	MYCONST6            //10
// )

/* [END] IOTA USAGE */

// func main() {
// 	// fmt.Printf("MYCONST1: [%d]\nMYCONST2: [%d]\nMYCONST3: [%d]\nMYCONST4: [%d]\nMYCONST5: [%d]\nMYCONST6: [%d]\n",
// 	// 	MYCONST1, MYCONST2, MYCONST3, MYCONST4, MYCONST5, MYCONST6)

// 	var nums = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

// 	var numsSlice []int = nums[2:5]

// 	numsSlice = append(numsSlice, 6, 7, 8, 9, 0, 1, 2)

// 	numsSlice[2] = 0

// 	fmt.Println("nums Length: ", len(nums), " Contents: ", nums, " Capacity: ", cap(nums))
// 	fmt.Println("numsSlice Length: ", len(numsSlice), " Contents: ", numsSlice, " Capacity: ", cap(numsSlice))

// }

// 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, ?, ?, ?, ?

func main() {
	// fizzbuzz()
	letterCount()
}

/**
 * Oct 27, 2022
 */
func fizzbuzz() {
	fmt.Print("Enter expected max loop: ")
	var maxLoop int
	fmt.Scanf("%d", &maxLoop)

	for i := 1; i <= maxLoop; i++ {
		fizz := "fizz"
		buzz := "buzz"

		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, fizz+buzz)
		} else if i%3 == 0 {
			fmt.Println(i, fizz)
		} else if i%5 == 0 {
			fmt.Println(i, buzz)
		} else {
			fmt.Println(i)
		}
	}
}

/**
 * Oct 27, 2022
 */
func letterCount() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("god says: ")
	scanner.Scan()

	words := scanner.Text()
	fmt.Printf("then let it be \"%s\"\n", words)

	var godLetters = make(map[string]int)

	for _, letter := range words {
		// fmt.Println("god's letter: ", string(letter))
		if _, letterExists := godLetters[string(letter)]; letterExists {
			godLetters[string(letter)]++
		} else {
			godLetters[string(letter)] = 1
		}
	}
	fmt.Println("so the world given holy letter", godLetters)
}
