package main

import(
	"fmt"
	"os"
	"log"
	"encoding/csv"
	"io"
	"strconv"
)

type ZipStats struct {
	TotalNumber int
	TotalNumberByType map[string]int
	HighestPopulation int
}

func main() {
	file, err := os.Open("./zip_code_database.csv")
	if err != nil {
		log.Fatal(err)
	}
	
	Total := 0
	TypeTotal := make(map[string]int)
	MaxPop := 0
	
	r := csv.NewReader(file)
	// for i := 1; i < r.length; i++ {
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		
		Total++
		TypeTotal[record[1]]++
		// if strconv.Atoi(record[14]) > MaxPop {
		// 	MaxPop = strconv.Atoi(record[14])
		// }
		
		// fmt.Println(record[0], record[1], record[14])
	}
	
	stats := ZipStats{
		TotalNumber: Total,
		TotalNumberByType: TypeTotal,
		// HighestPopulation: 98237
	}
	
	fmt.Println(stats)
	
	file.Close()
}