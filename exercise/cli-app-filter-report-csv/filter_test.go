package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFilterInputValidation(t *testing.T) {

	tests := []struct {
		input    []string
		expected string
	}{
		{
			input:    []string{"./string", "00-00-00", "00-00-00"},
			expected: "failed parse startDate",
		},
		{
			input:    []string{"./string", "2023-04-22T18:40:35+07:00", "00-00-00"},
			expected: "failed parse endDate",
		},
		{
			input:    []string{"./string", "2023-04-22T18:40:35+07:00", "2023-03-22T18:40:35+07:00"},
			expected: "invalid date time range",
		},
		{
			input:    []string{"./string", "2023-04-22T18:40:35+07:00", "2023-04-22T18:40:35+07:00"},
			expected: "failed to list file from directory",
		},
	}

	for _, test := range tests {
		err := filter(test.input[0], test.input[1], test.input[2])
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, test.expected)

	}
}

func TestSearchTransactionByDate(t *testing.T) {

	tests := []struct {
		transactions   []Transaction
		targetDateTime time.Time
		expected       int
	}{
		{
			transactions: []Transaction{
				{
					date: time.Date(2023, time.March, 20, 0, 0, 0, 0, time.Local),
				},
				{
					date: time.Date(2023, time.March, 21, 0, 0, 0, 0, time.Local),
				},
				{
					date: time.Date(2023, time.April, 21, 0, 0, 0, 0, time.Local),
				},
			},
			targetDateTime: time.Date(2023, time.April, 21, 0, 0, 0, 0, time.Local),
			expected:       2,
		},
		{
			transactions: []Transaction{
				{
					date: time.Date(2023, time.March, 20, 0, 0, 0, 0, time.Local),
				},
				{
					date: time.Date(2023, time.March, 21, 0, 0, 0, 0, time.Local),
				},
				{
					date: time.Date(2023, time.April, 21, 0, 0, 0, 0, time.Local),
				},
			},
			targetDateTime: time.Date(2023, time.February, 21, 0, 0, 0, 0, time.Local),
			expected:       -1,
		},
		{
			transactions: []Transaction{
				{
					date: time.Date(2023, time.March, 20, 0, 0, 0, 0, time.Local),
				},
				{
					date: time.Date(2023, time.March, 21, 0, 0, 0, 0, time.Local),
				},
				{
					date: time.Date(2023, time.April, 21, 0, 0, 0, 0, time.Local),
				},
			},
			targetDateTime: time.Date(2023, time.September, 21, 0, 0, 0, 0, time.Local),
			expected:       -1,
		},
		{
			transactions: []Transaction{
				{
					date: time.Date(2023, time.March, 20, 0, 0, 0, 0, time.Local),
				},
				{
					date: time.Date(2023, time.March, 22, 0, 0, 0, 0, time.Local),
				},
				{
					date: time.Date(2023, time.April, 21, 0, 0, 0, 0, time.Local),
				},
			},
			targetDateTime: time.Date(2023, time.March, 21, 11, 0, 0, 0, time.Local),
			expected:       1,
		},
	}

	for _, test := range tests {
		index := searchTransactionByDate(test.transactions, test.targetDateTime)
		assert.Equal(t, test.expected, index)
	}
}
