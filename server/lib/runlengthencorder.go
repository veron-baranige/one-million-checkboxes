package lib

import (
	"strconv"
	"strings"
)

type RunLengthEncoder struct {
	data map[uint32]bool
}

func NewRunLengthEncoder(checkboxMap map[uint32]bool) *RunLengthEncoder {
	return &RunLengthEncoder{
		data: checkboxMap,
	}
}

func (rle *RunLengthEncoder) GetEncodedValue() string {
	var runLength uint
	var currentValue bool
	var encodedValue strings.Builder

	for i := 1; i <= len(rle.data); i++ {
		value := rle.data[uint32(i)]
		if i == 1 {
			currentValue = value
			runLength = 1
		} else if currentValue != value {
			encodedValue.WriteString(getFormattedRunLength(currentValue, runLength) + ",")
			currentValue = value
			runLength = 1
		} else {
			runLength++
		}

		if i == len(rle.data) {
			encodedValue.WriteString(getFormattedRunLength(currentValue, runLength))
		}
	}
	return strings.TrimSuffix(encodedValue.String(), ",")
}

func getFormattedRunLength(value bool, length uint) string {
	valueStr := "F"
	if value {
		valueStr = "T"
	}
	return valueStr + strconv.Itoa(int(length))
}
