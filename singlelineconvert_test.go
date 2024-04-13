package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestConvert(t *testing.T) {
	csvData := `value,income,age,rooms,bedrooms,pop,hh
1e+05,8.3252,41,880,129,322,126
358500,8.3014,21,7099,1106,2401,1138`

	expectedJsonData := `{"value":100000,"income":8.3252,"age":41,"rooms":880,"bedrooms":129,"pop":322,"hh":126}
{"value":358500,"income":8.3014,"age":21,"rooms":7099,"bedrooms":1106,"pop":2401,"hh":1138}
`

	csvReader := strings.NewReader(csvData)
	jsonWriter := &bytes.Buffer{}

	err := convert(csvReader, jsonWriter)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if jsonWriter.String() != expectedJsonData {
		t.Errorf("Expected %s, got %s", expectedJsonData, jsonWriter.String())
	}
}
