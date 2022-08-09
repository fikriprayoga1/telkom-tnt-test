package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type ModelMoney struct {
	MoneyValue string `json:"money"`
	Total      int    `json:"total"`
}

func main() {
	var modelMoney []ModelMoney
	var err error
	var text string
	var value int
	var mTotal int
	var finalValue []byte

	log.Println("Please type input")
	fmt.Scanln(&text)

	value, err = strconv.Atoi(text)

	if err != nil {
		log.Print(err)
		log.Println("Please type number only")
		return
	}

	if value >= 100000 {
		mTotal = value / 100000
		value = value - (mTotal * 100000)
		modelMoney = append(modelMoney, ModelMoney{MoneyValue: "Rp. 100.000", Total: mTotal})
	}

	if value >= 50000 {
		mTotal = value / 50000
		value = value - (mTotal * 50000)
		modelMoney = append(modelMoney, ModelMoney{MoneyValue: "Rp. 50.000", Total: mTotal})
	}

	if value >= 20000 {
		mTotal = value / 20000
		value = value - (mTotal * 20000)
		modelMoney = append(modelMoney, ModelMoney{MoneyValue: "Rp. 20.000", Total: mTotal})
	}

	if value >= 10000 {
		log.Print(value)
		mTotal = value / 10000
		value = value - (mTotal * 10000)
		modelMoney = append(modelMoney, ModelMoney{MoneyValue: "Rp. 10.000", Total: mTotal})
	}

	if value >= 5000 {
		mTotal = value / 5000
		value = value - (mTotal * 5000)
		modelMoney = append(modelMoney, ModelMoney{MoneyValue: "Rp. 5.000", Total: mTotal})
	}

	if value >= 2000 {
		mTotal = value / 2000
		value = value - (mTotal * 2000)
		modelMoney = append(modelMoney, ModelMoney{MoneyValue: "Rp. 2.000", Total: mTotal})
	}

	if value >= 1000 {
		mTotal = value / 1000
		value = value - (mTotal * 1000)
		modelMoney = append(modelMoney, ModelMoney{MoneyValue: "Rp. 1.000", Total: mTotal})
	}

	if value >= 500 {
		mTotal = value / 500
		value = value - (mTotal * 500)
		modelMoney = append(modelMoney, ModelMoney{MoneyValue: "Rp. 500", Total: mTotal})
	}

	if value >= 200 {
		mTotal = value / 200
		value = value - (mTotal * 200)
		modelMoney = append(modelMoney, ModelMoney{MoneyValue: "Rp. 200", Total: mTotal})
	}

	if value >= 100 {
		mTotal = value / 100
		value = value - (mTotal * 100)
		modelMoney = append(modelMoney, ModelMoney{MoneyValue: "Rp. 100", Total: mTotal})
	}

	if (value < 100) && (value > 0) {
		value = 0
		modelMoney = append(modelMoney, ModelMoney{MoneyValue: "Rp. 100", Total: 1})
	}

	finalValue, err = json.MarshalIndent(modelMoney, "", "  ")
	if err != nil {
		log.Println("Converting object to json failed")
		return
	}

	log.Println(string(finalValue))

}
