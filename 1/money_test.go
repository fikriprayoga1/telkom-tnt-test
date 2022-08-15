package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoney1(t *testing.T) {
	var modelMoney []ModelMoney
	var finalValue []byte
	var result string
	var err error

	modelMoney = append(modelMoney, ModelMoney{MoneyValue: "Rp. 100.000", Total: 1})
	modelMoney = append(modelMoney, ModelMoney{MoneyValue: "Rp. 20.000", Total: 2})
	modelMoney = append(modelMoney, ModelMoney{MoneyValue: "Rp. 5.000", Total: 1})
	finalValue, err = json.MarshalIndent(modelMoney, "", "  ")
	if err != nil {
		t.Error("Converting object to json failed")

	}
	result = moneyReturn("145000")
	assert.Equal(t, string(finalValue), result)

}

func TestMoney2(t *testing.T) {
	var modelMoney []ModelMoney
	var finalValue []byte
	var result string
	var err error

	modelMoney = append(modelMoney, ModelMoney{MoneyValue: "Rp. 2.000", Total: 1})
	modelMoney = append(modelMoney, ModelMoney{MoneyValue: "Rp. 100", Total: 1})
	finalValue, err = json.MarshalIndent(modelMoney, "", "  ")
	if err != nil {
		t.Error("Converting object to json failed")

	}
	result = moneyReturn("2050")
	assert.Equal(t, string(finalValue), result)

}
