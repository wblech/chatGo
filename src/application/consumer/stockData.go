package consumer

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GetStockData(msg string) string {
	url := fmt.Sprintf("https://stooq.com/q/l/?s=%s&f=sd2t2ohlcv&h&e=csv", msg)
	_ = DownloadCSV(url, "stock.csv")
	msgToSend := ReadCSV("stock.csv")
	_ = os.Remove("stock.csv")
	return msgToSend
}

func DownloadCSV(url string, filename string) error {
	out, _ := os.Create(filename)
	defer out.Close()

	resp, err := http.Get(url)
	defer resp.Body.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		log.Fatal(err)
	}
	return err
}

func ReadCSV(filename string) string {
	csvFile, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Println(err)
	}
	var returnMSG string
	for i, line := range csvLines {
		if i == 0 {
			continue
		}
		returnMSG = fmt.Sprintf("%s %s", line[0], line[3])
	}
	return returnMSG
}
