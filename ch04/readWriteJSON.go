package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Value1  int `json:",omitempty"`
	Value2  int `json:",omitempty"`
	Value3  int `json:",omitempty"`
	Value4  int `json:",omitempty"`
	Value5  int `json:",omitempty"`
	Value6  int `json:",omitempty"`
	Value7  int `json:",omitempty"`
	Value8  int `json:",omitempty"`
	Value9  int `json:",omitempty"`
	Value10 int `json:",omitempty"`
}

func readJSON(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}

	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}
	in.Close()
	return nil
}

func changeValues(*s Record) error {
	for str, val := range s{
		fmt.Println(str, val) 
	}
	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a json file name")
		return
	}

	filename := arguments[1]
	var myStruct Record
	err := readJSON(filename, &myStruct)
	if err == nil {
		fmt.Println(myStruct.Value10)
	} else {
		fmt.Println(err)
	}

}
