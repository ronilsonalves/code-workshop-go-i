package main

import (
	"fmt"
	"internal/tickets/tickets.go/internal/tickets"
)

const (
	morning   = "MORNING"
	afternoon = "AFTERNOON"
	night     = "NIGHT"
	daybreak  = "DAYBREAK"
)

func main() {
	MainMenu()
}

func MainMenu() {
	option := 99
Loop:
	for {
		fmt.Println("============================================================")
		fmt.Println("MAIN MENU")
		fmt.Println("============================================================")
		fmt.Printf(
			"1 - Total trips to a specified destination\n" +
				"2 - Total trips by interval time\n" +
				"3 - Average number of trips per country dest.\n" +
				"0 - Exit program\n")
		fmt.Println("============================================================")
		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println(err.Error())
		}

		switch option {
		case 0:
			fmt.Println("...exiting")
			break Loop
		case 1:
			SubMenuGetTotalTickets()
		case 2:
			SubMenuGetCountByInterval()
		case 3:
			SubMenuAverageDestination()
		default:
			MainMenu()
		}
	}

}

func SubMenuGetTotalTickets() {
	destination := "NONE"
	fmt.Println("============================================================")
	fmt.Println("Insert the country name of destination")
	_, err := fmt.Scan(&destination)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		total, err := tickets.GetTotalTickets(destination)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("# peoples who traveled to %s: %d\n", destination, total)
		}
	}
}

func SubMenuGetCountByInterval() {
	option := 99
Loop:
	for {
		fmt.Println("============================================================")
		fmt.Println("Time interval menu")
		fmt.Println("============================================================")
		fmt.Printf(
			"1 - MORNING\n" +
				"2 - AFTERNOON\n" +
				"3 - NIGHT\n" +
				"4 - DAYBREAK\n" +
				"0 - Press 0 to return to previous menu\n")
		fmt.Println("============================================================")
		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println(err.Error())
		}

		switch option {
		case 0:
			fmt.Println("...exiting")
			break Loop
		case 1:
			fmt.Println("------------------------------------------------------------")
			total, err := tickets.GetCountByInterval(morning)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("## \tnumber of trips during the %s: %d\t\t ##\n", morning, total)
			}
			break
		case 2:
			fmt.Println("------------------------------------------------------------")
			total, err := tickets.GetCountByInterval(afternoon)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("## \tnumber of trips during the %s: %d\t\t ##\n", afternoon, total)
			}
			break
		case 3:
			fmt.Println("------------------------------------------------------------")
			total, err := tickets.GetCountByInterval(night)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("## \tnumber of trips during the %s: %d\t\t ##\n", night, total)
			}
			break
		case 4:
			fmt.Println("------------------------------------------------------------")
			total, err := tickets.GetCountByInterval(daybreak)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("## \tnumber of trips during the %s: %d\t\t ##\n", daybreak, total)
			}
			break
		default:
			SubMenuGetCountByInterval()
		}
	}
}

func SubMenuAverageDestination() {

	total, err := tickets.AverageTicketPerCountryDestination()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("------------------------------------------------------------")
		fmt.Printf("## \tAverage of tickets selled per country dest.: %.2f\t##\n", total)
	}
}
