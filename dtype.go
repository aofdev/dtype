package dtype

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
)

// DefaultStringWithNullable if the value equal is string empty returns raw sting and nil pointer
func DefaultStringWithNullable(inputString string) (string, *string) {
	var nullableStrType *string
	nullableStrType = nil
	if inputString != "" {
		nullableStrType = &inputString
	}
	return inputString, nullableStrType
}

// DefaultIntWithNullable if the value equal is 0 returns raw int and nil pointer
func DefaultIntWithNullable(inputInt int) (int, *int) {
	var nullableIntType *int
	nullableIntType = nil
	if inputInt != 0 {
		nullableIntType = &inputInt
	}
	return inputInt, nullableIntType
}

// DefaultFloat64WithNullable if the value equal is 0.0 returns raw int and nil pointer
func DefaultFloat64WithNullable(inputFloat float64) (float64, *float64) {
	var nullableFloat64Type *float64
	nullableFloat64Type = nil
	if inputFloat != 0.0 {
		nullableFloat64Type = &inputFloat
	}
	return inputFloat, nullableFloat64Type
}

// DefaultArrayStringWithNullable if the value equal is array length 0 returns raw array and nil pointer
func DefaultArrayStringWithNullable(inputArray []string) ([]string, *[]string) {
	var nullableArrayType *[]string
	nullableArrayType = nil
	if len(inputArray) != 0 {
		nullableArrayType = &inputArray
	}
	return inputArray, nullableArrayType
}

// DefaultArrayIntWithNullable if the value equal is array length 0 returns raw array and nil pointer
func DefaultArrayIntWithNullable(inputArray []int) ([]int, *[]int) {
	var nullableArrayType *[]int
	nullableArrayType = nil
	if len(inputArray) != 0 {
		nullableArrayType = &inputArray
	}
	return inputArray, nullableArrayType
}

// DefaultBooleanWithNullable if the value equal is true returns raw boolean and nil pointer
func DefaultBooleanWithNullable(inputBoolean bool) (bool, *bool) {
	var nullableBooleanType *bool
	nullableBooleanType = nil
	if inputBoolean {
		nullableBooleanType = &inputBoolean
	}
	return inputBoolean, nullableBooleanType
}

// DefaultJSONUnstructuredWithNullable if have key returns raw json and nil pointer
func DefaultJSONUnstructuredWithNullable(inputJSON, key string) (string, *string) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var nullableStringType *string
	result := gjson.Get(inputJSON, key)
	if !result.Exists() {
		var newMap map[string]interface{}
		err := json.Unmarshal([]byte(inputJSON), &newMap)
		if err != nil {
			panic(fmt.Errorf("Error is json unmarshal: %s", err))
		}
		newMap[key] = nil
		newData, _ := json.Marshal(newMap)
		inputJSON = string(newData)
	} else {
		nullableStringType = &key
	}
	return inputJSON, nullableStringType
}
