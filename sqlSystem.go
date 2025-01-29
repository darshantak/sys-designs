package main

import (
	"fmt"
	"reflect"
	"time"
)

type Database struct {
	Tables map[string]Table
}
type Column struct {
	Name string
	Type reflect.Type
}
type Table struct {
	Name      string
	Columns   []Column
	Data      [][]interface{}
	CreatedOn time.Time
}

func newDatabase() *Database {
	return &Database{
		Tables: make(map[string]Table),
	}
}

// func (db *Database) addTable(name string, cols []string, createdOn time.Time) {
func (db *Database) addTable(table Table) {
	for _, t := range db.Tables {
		if t.Name == table.Name {
			fmt.Println("Table already exists")
			return
		}
	}
	db.Tables[table.Name] = table
	fmt.Printf("Table %s added successfull \n", table.Name)
}
func (db *Database) insertData(name string, data []interface{}) {
	// if _, exists := db.Tables[name]; !exists {
	// 	return
	// }
	table, exists := db.Tables[name]
	if !exists {
		fmt.Println("This table doesn't exist")
		return
	}

	if len(data) != len(table.Columns) {
		fmt.Println("Data is not uniform with the columns")
		return
	}

	for i, value := range data {
		if reflect.TypeOf(value) != table.Columns[i].Type {
			fmt.Println("Types not matching. Please check the input")
			return
		}
	}

	table.Data = append(table.Data, data)
	db.Tables[name] = table
	fmt.Println("Data is inserted")
}
func (db *Database) queryData(tName string, conditions map[string]interface{}) {
	// select * from tName where colnName = "alice"
	// {"colName": "alice"}

	t, exists := db.Tables[tName]
	if !exists {
		fmt.Println("This table doesn't exist")
		return
	}
	cols := make(map[string]int)
	for i, c := range t.Columns {
		cols[c.Name] = i
	}

	result := [][]interface{}{}
	for _, row := range t.Data {
		match := true
		for colName, value := range conditions {
			idx, exists := cols[colName]
			if !exists {
				fmt.Println("Column doesnot exist")
				return
			}
			if row[idx] != value {
				match = false
				break
			}
		}
		if match {
			result = append(result, row)
		}
	}
	if len(result) == 0 {
		fmt.Println("No matching record found")
		return
	}
	for _, row := range result {
		fmt.Println(row)
	}
}

func SqlSystem() {
	newDB := newDatabase()

	table1 := Table{
		Name: "table1",
		Columns: []Column{
			{Name: "col1", Type: reflect.TypeOf("int")},
			{Name: "col2", Type: reflect.TypeOf("string")}},
		CreatedOn: time.Now(),
	}
	table2 := Table{
		Name: "table2",
		Columns: []Column{
			{Name: "col1", Type: reflect.TypeOf("string")},
			{Name: "col2", Type: reflect.TypeOf("int")}},
		CreatedOn: time.Now(),
	}
	newDB.addTable(table1)
	newDB.addTable(table2)

	// Insert data
	newDB.insertData("table1", []interface{}{"Alice", "1"})
	newDB.insertData("table1", []interface{}{"Bob", "2"})
	newDB.insertData("table2", []interface{}{"Test", "101"})

	// Print database state
	fmt.Println(newDB)
}
