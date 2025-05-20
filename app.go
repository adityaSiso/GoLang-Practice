// this is a special package. Entry for the project/module.
package main

// Formatting Module

// func main() {
// 	var choice int
// 	var withdrawl_amount, deposite_amount float64
// 	account_balance := 1000.00

// 	utilities.ReadDataFromFile("<Final Output>")
// 	utilities.WriteDataToFile(1000.00)

// 	// `for {}` infinite loop
// 	// `for condition` infinite loop until condition evaluates to true.
// 	for i := 0; i < 200; i++ {
// 		welcome()
// 		fmt.Print("Your Choice: ")
// 		fmt.Scan(&choice)
// 		switch choice {
// 		case 1:
// 			// line 1
// 			// line 2
// 			// ...
// 		case 2:
// 			// ...
// 		default:
// 			// ...
// 		}
// 		if choice == 3 {

// 			fmt.Println("Account Balance:", account_balance)

// 		} else if choice == 1 {

// 			fmt.Print("Enter Withdral Amount: ")
// 			fmt.Scan(&withdrawl_amount)

// 			if withdrawl_amount > float64(account_balance) || account_balance <= 0 {

// 				fmt.Println("Error!!")
// 				// return
// 				// break
// 				continue

// 			} else {

// 				account_balance -= withdrawl_amount

// 			}

// 		} else if choice == 2 {

// 			fmt.Print("Enter Deposite Amount: ")
// 			fmt.Scan(&deposite_amount)

// 			if deposite_amount <= 0 {
// 				fmt.Println("Error!!")
// 				// return
// 				continue
// 			}
// 			account_balance += deposite_amount

// 		} else {
// 			fmt.Println("Goodbye!")
// 			break
// 		}
// 	}
// 	fmt.Println("Have a Nice Day :)")
// }

// func calculate() {
// 	// Cannot update `const`` once assigned while `var` can be re-assigned and modified.
// 	const T = "Constant"
// 	var a float64 = 10
// 	b := 20.0
// 	var c = a + b
// 	z, y := 1, 2
// 	fmt.Printf("%.2f", 100.234123)
// 	fmt.Println()
// 	fmt.Println(z, y)
// 	fmt.Println("Value of C:", c, ", Type of C:", reflect.TypeOf((c)))
// 	fmt.Printf(`First Line
// 	.
// 	.
// 	.
// 	Last Line.`)
// }

// Main function that will be executed by GO.
func main() {

	// start_time := time.Now()
	// var count int64
	// for i := 0; i <= int(math.Pow10(9)); i++ {
	// 	select {
	// 	default:
	// 		// pass
	// 	}
	// }
	// end_time := time.Now()
	// fmt.Println("Total time taken to run a one billion iteration loop",
	// 	end_time.Sub(start_time))

	// collection_main()
	map_practice()
	// pointerFunction()
	// var a, b, c int
	// fmt.Scan(&a, &b, &c)
	// fmt.Println("Hello World", " Go!", a, b, c)
	// calculate()
}
