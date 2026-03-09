package main

import (
	"fmt"
	"strconv"
)

func IsValidSlice(userInputSlice []string) (ItemId, itemName string, itemQuantity int, itemPrice float64, errorFalg error) {
	if len(userInputSlice) > 0 && len(userInputSlice) <= 5 {
		ItemId = userInputSlice[1]
		itemName = userInputSlice[2]
		itemQuantity, _ = strconv.Atoi(userInputSlice[3])
		itemPrice, _ = strconv.ParseFloat(userInputSlice[4], 64)
		return ItemId, itemName, itemQuantity, itemPrice, nil

	} else {
		errorFalg = fmt.Errorf("complete the command ")
		return
	}
}
