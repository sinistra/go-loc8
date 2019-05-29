package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func main() {
	fileName := path.Join("..", "data", "SR_HFL_PFL_Delta_CSV_20190409", "SR_SAT_ADDR_HFL_PFL_Delta_20190409_201407.csv")
	fmt.Println(fileName)

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = ','
	r.Comment = '#'

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}
