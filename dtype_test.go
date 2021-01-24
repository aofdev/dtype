package dtype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultStringWithNullableReturnRawString(t *testing.T) {
	input := "test str"
	want := "test str"
	got, err := DefaultStringWithNullable(input)
	assert.NotNil(t, err)
	assert.Equal(t, want, got)
}

func TestDefaultStringWithNullableReturnNull(t *testing.T) {
	input := ""
	got, err := DefaultStringWithNullable(input)
	assert.Nil(t, err)
	assert.Empty(t, got)
}

func TestDefaultIntWithNullableReturnRawInt(t *testing.T) {
	input := 214748364700000
	want := 214748364700000
	got, err := DefaultIntWithNullable(input)
	assert.NotNil(t, err)
	assert.Equal(t, want, got)
}

func TestDefaultIntWithNullableReturnNull(t *testing.T) {
	input := 0
	got, err := DefaultIntWithNullable(input)
	assert.Nil(t, err)
	assert.Empty(t, got)
}

func TestDefaultFloat64WithNullableReturnRawInt(t *testing.T) {
	input := 22111100.202424
	want := 22111100.202424
	got, err := DefaultFloat64WithNullable(input)
	assert.NotNil(t, err)
	assert.Equal(t, want, got)
}

func TestDefaultFloat64WithNullableReturnNull(t *testing.T) {
	input := 0.0
	got, err := DefaultFloat64WithNullable(input)
	assert.Nil(t, err)
	assert.Empty(t, got)
}

func TestDefaultArrayStringWithNullableReturnRawArrayString(t *testing.T) {
	input := []string{"a", "b", "c"}
	want := []string{"a", "b", "c"}
	got, err := DefaultArrayStringWithNullable(input)
	assert.NotNil(t, err)
	assert.Equal(t, want, got)
}

func TestDefaultArrayStringWithNullableReturnNull(t *testing.T) {
	input := []string{}
	got, err := DefaultArrayStringWithNullable(input)
	assert.Nil(t, err)
	assert.Empty(t, got)
}

func TestDefaultArrayIntWithNullableReturnRawArrayInt(t *testing.T) {
	input := []int{1, 2, 3, 4}
	want := []int{1, 2, 3, 4}
	got, err := DefaultArrayIntWithNullable(input)
	assert.NotNil(t, err)
	assert.Equal(t, want, got)
}

func TestDefaultArrayIntWithNullableReturnNull(t *testing.T) {
	input := []int{}
	got, err := DefaultArrayIntWithNullable(input)
	assert.Nil(t, err)
	assert.Empty(t, got)
}

func TestDefaultBooleanWithNullableReturnRawBoolean(t *testing.T) {
	input := bool(true)
	want := bool(true)
	got, err := DefaultBooleanWithNullable(input)
	assert.NotNil(t, err)
	assert.Equal(t, want, got)
}

func TestDefaultBooleanWithNullableReturnNull(t *testing.T) {
	input := bool(false)
	got, err := DefaultBooleanWithNullable(input)
	assert.Nil(t, err)
	assert.Empty(t, got)
}

func TestDefaultJSONUnstructuredWithNullableReturnRawJSON(t *testing.T) {
	inputJSON := `{"name":{"first":"Janet","last":"Prichard"},"age":47, "country": "TH"}`
	inputKey := "country"
	want := `{"name":{"first":"Janet","last":"Prichard"},"age":47, "country": "TH"}`
	got, err := DefaultJSONUnstructuredWithNullable(inputJSON, inputKey)
	assert.NotNil(t, err)
	assert.Equal(t, want, got)
}

func TestDefaultJSONUnstructuredWithNullableReturnNull(t *testing.T) {
	inputJSON := `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	inputKey := "country"
	want := `{"name":{"first":"Janet","last":"Prichard"},"age":47, "country": null}`
	got, err := DefaultJSONUnstructuredWithNullable(inputJSON, inputKey)
	assert.Nil(t, err)
	assert.JSONEq(t, want, got)
}
