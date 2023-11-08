# Filter Report CLI App
## Overview
CLI application written in Golang to filter the transaction data within the defined time range. The transaction data is separated into different files and ordered based on timestamps. 

## How To Use
Sample command to filter available filter inside folder [reports](./reports):
```bash
./filter.exe -d "./reports" -s "2023-04-17T20:40:35+07:00" -e "2023-05-05T10:00:13+07:00"
```
The sample command will run action transaction filter based on given time range, inclusive start date and exclusive end date, so if you try to filter using the same start date and end date it will return empty result even though there is a row. Then, if successful, a file named `result.csv` will be generated. Filter cli app only read file report with postfix `_report` and extension file `.csv`

Filter cli app has the following options:
| Options  | Description | Default Value |
| ------------- | ------------- | ------------- |
| -d  | directory, put correct directory name of reports and only read file report with postfix `_report` and extension file `.csv` e.g. `100_report.csv`| ./ |
| -s  | start date, format RFC3339  | date time today |
| -e  | end date, format RFC3339  | date time today + 1 day |

## Unit Test
Filter cli app build with unit test in it to test specific function of the app.
Command to run unit test:
```bash
go test filter -v
```

## Dependencies
* Unit Test: [Testify](https://github.com/stretchr/testify)