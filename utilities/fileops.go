package utilities

import (
	"fmt"
	"os"
	"strconv"
)

func WriteDataToFile(data float64) {
	dataToText := fmt.Sprint(data)
	os.WriteFile("balance.txt", []byte(dataToText), 0644)
}

// Notice the Case of function name.
func ReadDataFromFile(file string) (float64, error) {
	data, err := os.ReadFile(file)

	if err != nil {
		panic("Error in reading file...")
	}
	Stringdata := string(data)
	return strconv.ParseFloat(Stringdata, 64)
}
