package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Transaction struct {
	no     int
	date   time.Time
	detail string
	amount string
}

func main() {
	var (
		directory string
		end       string
		start     string
	)
	flag.CommandLine.StringVar(&directory, "d", "./reports", "put directory of reports")
	flag.CommandLine.StringVar(&start, "s", time.Now().UTC().Format(time.RFC3339), "start time")
	flag.CommandLine.StringVar(&end, "e", time.Now().Add(time.Hour*24).UTC().Format(time.RFC3339), "end time")
	flag.Parse()

	err := filter(directory, start, end)
	if err != nil {
		log.Fatalf("Unable to filter report, err: %v", err)
	}

	log.Println("Success filter report!")

}

func filter(dir, start, end string) error {
	startDate, err := time.Parse(time.RFC3339, start)
	if err != nil {
		return fmt.Errorf("failed parse startDate to RFC3339, %v", err)
	}

	endDate, err := time.Parse(time.RFC3339, end)
	if err != nil {
		return fmt.Errorf("failed parse endDate to RFC3339, %v", err)
	}

	if startDate.After(endDate) {
		return errors.New("invalid date time range")
	}

	transactions, err := getTransactions(dir, startDate, endDate)
	if err != nil {
		return err
	}

	err = writeToCSV(transactions)
	if err != nil {
		return err
	}

	return nil

}

func getSortedFiles(dir string) (result []string, err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf(`failed to list file from directory "%s", error %v`, dir, err)
	}

	ext := "_report.csv"
	var numberOfFiles []int
	for _, v := range files {
		if strings.Contains(v.Name(), ext) {
			r := strings.Split(v.Name(), "_")
			c, _ := strconv.Atoi(r[0])
			numberOfFiles = append(numberOfFiles, c)
		}
	}
	sort.Ints(numberOfFiles)

	if numberOfFiles == nil {
		return nil, errors.New("can't find any report file")
	}

	for _, n := range numberOfFiles {
		fileName := fmt.Sprintf("%d%s", n, ext)
		result = append(result, fileName)
	}
	return
}

func getTransactions(dir string, startDate, endDate time.Time) ([]Transaction, error) {
	files, err := getSortedFiles(dir)
	if err != nil {
		return nil, err
	}

	transactions := make([]Transaction, 0)
	var firstIndex, lastIndex int = -1, -1

	for _, file := range files {
		filePath := dir + "/" + file
		result, err := readReportFile(filePath)
		if err != nil {
			return nil, fmt.Errorf(`failed to read file "%s", error %v`, filePath, err)
		}

		if firstIndex == -1 {
			firstIndex = searchTransactionByDate(result, startDate)
		}

		if lastIndex == -1 {
			lastIndex = searchTransactionByDate(result, endDate)
		}

		if firstIndex != -1 && lastIndex != -1 {
			fmt.Println("First Index", firstIndex)
			transactions = append(transactions, result[firstIndex:lastIndex]...)
			break
		} else if firstIndex != -1 && lastIndex == -1 {
			transactions = append(transactions, result[firstIndex:]...)
			firstIndex = 0
			continue
		} else {
			break
		}
	}
	return transactions, nil
}

func searchTransactionByDate(data []Transaction, target time.Time) (index int) {
	indexFirst := 0
	indexLast := len(data) - 1
	result := -1

	if data[indexFirst].date.After(target) || data[indexLast].date.Before(target) {
		return -1
	}

	for indexFirst <= indexLast {
		mid := (indexLast + indexFirst) / 2

		if data[mid].date.Equal(target) {
			return mid
		}
		if target.Before(data[mid].date) {
			indexLast = mid - 1
			result = indexFirst
		} else {
			indexFirst = mid + 1
			result = indexLast + 1
		}
	}

	return result
}

func readReportFile(path string) ([]Transaction, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvLines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	var transactions []Transaction
	for _, v := range csvLines {
		transaction, err := parse(v)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func parse(data []string) (Transaction, error) {
	var result Transaction
	no, err := strconv.Atoi(data[0])
	if err != nil {
		return result, err
	}
	result.no = no
	date, err := time.Parse(time.RFC3339, data[1])
	if err != nil {
		return result, err
	}
	result.date = date
	result.detail = data[2]
	result.amount = data[3]
	return result, nil
}

func writeToCSV(data []Transaction) error {
	file, err := os.Create("./result.csv")
	if err != nil {
		return fmt.Errorf("failed to create file result.csv, err: %v", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	var result [][]string
	for _, v := range data {
		row := []string{strconv.Itoa(v.no), v.date.Format(time.RFC3339), v.detail, v.amount}
		result = append(result, row)
	}

	err = w.WriteAll(result)
	if err != nil {
		return fmt.Errorf("failed to write all data into result.csv, err: %v", err)
	}

	return nil
}
