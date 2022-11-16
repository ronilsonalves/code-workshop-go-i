package tickets

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	morning   = "MORNING"
	afternoon = "AFTERNOON"
	night     = "NIGHT"
	daybreak  = "DAYBREAK"
)

type Ticket struct {
	ID          int
	Passenger   string
	Email       string
	Destination string
	Time        string
	Price       float64
}

func (t *Ticket) Set(id int, passenger string, email string, destination string, time string, price float64) {
	t.ID = id
	t.Passenger = passenger
	t.Email = email
	t.Destination = destination
	t.Time = time
	t.Price = price
}

var fileToRead, _ = filepath.Abs("./tickets.csv")

// GetTotalTickets by destination /*
func GetTotalTickets(destination string) (int, error) {
	var total int

	records, err := OpenAndReadFile()
	if err != nil {
		return total, fmt.Errorf("an error ocurred while trying to open/read the file %s", fileToRead)
	}

	// Calculate how many people travelled to specified destination
	for _, record := range records {
		if strings.Contains(record.Destination, destination) {
			total = total + 1
		}
	}

	return total, nil
}

// GetCountByInterval - Return the total trips made during the specified time interval /*
func GetCountByInterval(time string) (int, error) {
	var total int

	switch time {
	case morning:
		return GetMornings(), nil
	case afternoon:
		return GetAfternoons(), nil
	case night:
		return GetNights(), nil
	case daybreak:
		return GetDaybreaks(), nil
	default:
		return total, fmt.Errorf("an invalid time interval selected, please verify the data insertion and try again")
	}
}

func GetMornings() int {
	var total int

	records, err := OpenAndReadFile()
	if err != nil {
		return 0
	}

	for _, record := range records {
		if strings.HasPrefix(record.Time, "7:") || strings.HasPrefix(record.Time, "8:") || strings.HasPrefix(record.Time, "9:") || strings.HasPrefix(record.Time, "10:") || strings.HasPrefix(record.Time, "11:") || strings.HasPrefix(record.Time, "12:") {
			total = total + 1
		}
	}

	return total
}

func GetAfternoons() int {
	var total int

	records, err := OpenAndReadFile()
	if err != nil {
		return 0
	}

	for _, record := range records {
		if strings.HasPrefix(record.Time, "13:") || strings.HasPrefix(record.Time, "14:") || strings.HasPrefix(record.Time, "15:") || strings.HasPrefix(record.Time, "16:") || strings.HasPrefix(record.Time, "17:") || strings.HasPrefix(record.Time, "18:") || strings.HasPrefix(record.Time, "19:") {
			total = total + 1
		}
	}

	return total
}

func GetNights() int {
	var total int

	records, err := OpenAndReadFile()
	if err != nil {
		return 0
	}

	for _, record := range records {
		if strings.HasPrefix(record.Time, "20:") || strings.HasPrefix(record.Time, "21:") || strings.HasPrefix(record.Time, "22:") || strings.HasPrefix(record.Time, "23:") {
			total = total + 1
		}
	}

	return total
}

func GetDaybreaks() int {
	var total int

	records, err := OpenAndReadFile()
	if err != nil {
		return 0
	}

	for _, record := range records {
		if strings.HasPrefix(record.Time, "0:") || strings.HasPrefix(record.Time, "1:") || strings.HasPrefix(record.Time, "2:") || strings.HasPrefix(record.Time, "3:") || strings.HasPrefix(record.Time, "4:") || strings.HasPrefix(record.Time, "5:") || strings.HasPrefix(record.Time, "6:") {
			total = total + 1
		}
	}

	return total
}

// AverageTicketPerCountryDestination - Return an average of tickets sold by country destination/*
func AverageTicketPerCountryDestination() (float64, error) {

	var countries []string

	records, err := OpenAndReadFile()
	if err != nil {
		return 0, err
	}

	keys := make(map[string]bool)

	for _, record := range records {
		if _, value := keys[record.Destination]; !value {
			keys[record.Destination] = true
			countries = append(countries, record.Destination)
		}
	}

	return float64(len(records)) / float64(len(countries)), nil
}

func OpenAndReadFile() ([]Ticket, error) {
	var tickets []Ticket
	res, err := os.Open(fileToRead)
	if err != nil {
		return nil, err
	}

	read := csv.NewReader(res)
	read.Comma = ','

	records, err := read.ReadAll()
	if err != nil {
		return nil, err
	}

	// Converter from an array of strings into a slice of Tickets
	// instead of return a [][]string
	for _, record := range records {
		id, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			return []Ticket{}, err
		}
		passenger := record[1]
		email := record[2]
		destination := record[3]
		time := record[4]
		price, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			return []Ticket{}, err
		}

		ticket := Ticket{}

		ticket.Set(int(id), passenger, email, destination, time, price)

		tickets = append(tickets, ticket)
	}

	err = res.Close()
	if err != nil {
		return nil, err
	}

	return tickets, nil
}
