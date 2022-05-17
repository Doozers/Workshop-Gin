package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	s "workshop/structure"
)

func GetDatabase() ([]s.Item, error) {
	items := new([]s.Item)

	jsonFile, err := os.Open("database/database.json")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &items)
	if err != nil {
		return nil, err
	}

	fmt.Println(items)
	return *items, nil
}

func AddItemInDB(i s.Item) error {
	var items []s.Item

	jsonFile, err := os.Open("database/database.json")
	if err != nil {
		fmt.Println(err)
		return err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &items)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, item := range items {
		if item.Name == i.Name {
			return fmt.Errorf("item already exist")
		}
	}

	items = append(items, i)
	fmt.Println(items)

	file, err := json.MarshalIndent(items, "", " ")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = ioutil.WriteFile("database/database.json", file, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func DeleteItemInDB(i s.Item) error {
	var items []s.Item

	jsonFile, err := os.Open("database/database.json")
	if err != nil {
		fmt.Println(err)
		return err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &items)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for index, item := range items {
		if item.Name == i.Name {
			items = append(items[:index], items[index+1:]...)
		}
	}

	file, err := json.MarshalIndent(items, "", " ")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = ioutil.WriteFile("database/database.json", file, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func UpdateItemInDB(i s.Item) error {
	var items []s.Item

	jsonFile, err := os.Open("database/database.json")
	if err != nil {
		fmt.Println(err)
		return err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &items)
	if err != nil {
		fmt.Println(err)
		return err
	}

	inDB := false
	for index, item := range items {
		if item.Name == i.Name {
			items[index] = i
			inDB = true
		}
	}
	if !inDB {
		return fmt.Errorf("item not found")
	}

	file, err := json.MarshalIndent(items, "", " ")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = ioutil.WriteFile("database/database.json", file, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func AddItemsInDB(i []s.Item) error {
	var items []s.Item

	jsonFile, err := os.Open("database/database.json")
	if err != nil {
		fmt.Println(err)
		return err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &items)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, i := range i {
		for _, item := range items {
			if item.Name == i.Name {
				return fmt.Errorf("item: \"" + item.Name + "\" already exist")
			}
		}
	}

	items = append(items, i...)
	fmt.Println(items)
	file, err := json.MarshalIndent(items, "", " ")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = ioutil.WriteFile("database/database.json", file, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
