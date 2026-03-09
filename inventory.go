package main

import "fmt"

type Item struct {
	Price    float64
	Quantity int
	ItemId   string
	ItemName string
}

type ItemStore map[string]Item

func (I ItemStore) AddItem(Id, ItemName string, Quantity int, Price float64) map[string]Item {
	Item1 := Item{
		ItemId:   Id,
		ItemName: ItemName,
		Quantity: Quantity,
		Price:    Price,
	}
	I[Id] = Item1
	return I
}

func (I ItemStore) UpdateItem(Id, ItemName string, Quantity int, Price float64) {

	if I == nil {
		fmt.Print("cannot Item does not exist ")
		return
	}
	val, ok := I[Id]

	if !ok {
		fmt.Print("cannot Item does not exist ")
		return
	}
	val = Item{
		ItemId:   Id,
		ItemName: ItemName,
		Quantity: Quantity,
		Price:    Price,
	}
	I[Id] = val
}

func (I ItemStore) DeleteItem(Id string) {
	if I == nil {
		fmt.Print("cannot delete an empty map")
		return
	}
	_, ok := I[Id]

	if !ok {
		fmt.Print("cannot delete Item, Item does not exist ")
		return
	}
	delete(I, Id)
}
