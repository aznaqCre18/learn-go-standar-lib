package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	csvString := "aziz,nur,abdul,qodir\n" + "evi,sujiyati\n" + "husin,abdul,kadir"

	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"aziz", "nur", "abdul", "qodir"})
	_ = writer.Write([]string{"evi", "sujiyati"})
	_ = writer.Write([]string{"husin", "abdul", "kadir"})

	// memaksa data ditampilkan menggunakan Flush() method
	writer.Flush()

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
