package lab2

import "fmt"

func main() {
	const (
		COST_CLASS_A = 15
		COST_CLASS_B = 12
		COST_CLASS_C = 9
	)

	var classATicketsSold, classBTicketsSold, classCTicketsSold int
	fmt.Print("Enter the number of class A tickets sold:")
	fmt.Scan(&classATicketsSold)
	fmt.Print("Enter the number of class B tickets sold:")
	fmt.Scan(&classBTicketsSold)
	fmt.Print("Enter the number of class C tickets sold:")
	fmt.Scan(&classCTicketsSold)

	fmt.Printf("\nTotal income from sale of tickets is $%.2f\n", float64((classATicketsSold*COST_CLASS_A)+(classBTicketsSold*COST_CLASS_B)+(classCTicketsSold*COST_CLASS_C)))
}
