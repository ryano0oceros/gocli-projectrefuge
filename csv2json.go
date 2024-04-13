package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type House struct {
	Value    float64 `json:"value"`
	Income   float64 `json:"income"`
	Age      int     `json:"age"`
	Rooms    int     `json:"rooms"`
	Bedrooms int     `json:"bedrooms"`
	Pop      int     `json:"pop"`
	HH       int     `json:"hh"`
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: csvtojl <input.csv> <output.jl>")
		os.Exit(1)
	}

	csvFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	jsonFile, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	err = convert(csvFile, jsonFile)
	if err != nil {
		log.Fatal(err)
	}
}

func convert(csvFile io.Reader, jsonFile io.Writer) error {
	reader := csv.NewReader(csvFile)

	// Skip the header row.
	_, err := reader.Read()
	if err != nil {
		return err
	}

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records {
		value, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			return err
		}

		income, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return err
		}

		age, err := strconv.Atoi(record[2])
		if err != nil {
			return err
		}

		rooms, err := strconv.Atoi(record[3])
		if err != nil {
			return err
		}

		bedrooms, err := strconv.Atoi(record[4])
		if err != nil {
			return err
		}

		pop, err := strconv.Atoi(record[5])
		if err != nil {
			return err
		}

		hh, err := strconv.Atoi(record[6])
		if err != nil {
			return err
		}

		house := House{
			Value:    value,
			Income:   income,
			Age:      age,
			Rooms:    rooms,
			Bedrooms: bedrooms,
			Pop:      pop,
			HH:       hh,
		}

		jsonData, err := json.Marshal(house)
		if err != nil {
			return err
		}

		jsonFile.Write(jsonData)
		jsonFile.Write([]byte("\n"))
	}

	return nil
}
