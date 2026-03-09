package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run() {
	fmt.Println("Inventory Management System")
	fmt.Println("Commands: add, view, update, delete, exit")
	reader := bufio.NewReader(os.Stdin)
	var myMap ItemStore
	myMap = make(ItemStore)
	for {
		fmt.Print("> ")
		userInput, _ := reader.ReadString('\n')
		userInputSlice := strings.Fields(userInput)
		if len(userInputSlice) == 0 {
			fmt.Print("enter a valid user command \n")
			continue
		}
		command := userInputSlice[0]

		if len(userInputSlice) > 0 && len(userInputSlice) <= 5 {
			switch command {
			case "add":
				if len(userInputSlice) == 5 {
					itemId, itemName, itemQuantity, itemPrice, err := IsValidSlice(userInputSlice)
					if err != nil {
						fmt.Println(err)
					}
					myMap.AddItem(itemId, itemName, itemQuantity, itemPrice)
					val, _ := myMap[itemId]
					fmt.Printf("Added: ItemId:%s ItemName:%s ItemQuantity:%v ItemPrice:%v \n", val.ItemId, val.ItemName, val.Quantity, val.Price)
					continue
				}
			case "update":
				itemId, itemName, itemQuantity, itemPrice, err := IsValidSlice(userInputSlice)
				if err != nil {
					fmt.Println(err)
				}
				myMap.UpdateItem(itemId, itemName, itemQuantity, itemPrice)
				val, _ := myMap[itemId]
				fmt.Printf("Updated: ItemId:%s ItemName:%s ItemQuantity:%v ItemPrice:%v \n", val.ItemId, val.ItemName, val.Quantity, val.Price)
				continue
			case "delete":
				if len(userInputSlice) == 1 {
					itemId := userInputSlice[1]
					fmt.Println(itemId)
					val, _ := myMap[itemId]
					fmt.Printf("Deleted: ItemId:%s ItemName:%s\n", val.ItemId, val.ItemName)
					myMap.DeleteItem(itemId)
					continue
				} else {
					fmt.Println("delete command only takes itemId")
				continue
				}
			case "view":
				fmt.Println("-- Current Inventory ---")
				for _, val := range myMap {
					fmt.Printf("ID:%s ItemName:%s ItemQuantity:%v ItemPrice:%v \n", val.ItemId, val.ItemName, val.Quantity, val.Price)
				}
				fmt.Println("------------------------")
			case "exit":
				fmt.Println("thanks for using the inventory goodbye ")
				os.Exit(1)
			}
		} else {
			fmt.Print("enter a valid user command \n")

		}
	}
}
