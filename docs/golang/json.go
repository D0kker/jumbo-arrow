package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// only exportable fields are accessible for unmarshal
type greeting struct {
	Message      string
	OtherMessage string `json:"otherMessage"`
}

type person struct {
	Lastname  string  `json:"lname"`
	Firstname string  `json:"fname"`
	Address   address `json:"address"`
}
type address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zipcode"`
}

// with no tags, marshals with key name capitalized
type goodbye struct {
	SomeMessage string
}

type book struct {
	ISBN          string `json:"isbn"`
	Title         string `json:"title"`
	YearPublished int    `json:"yearpub"`
	Author        string `json:"author"`
	// omitempty tag to avoid zero value on uninitialized fields
	CoAuthor string `json:"coauthor,omitempty"`
}

// "keyName,omitempty" keyName is key in JSON, if field is zero, its omitted
// ",omitempty" key in JSON is field name
// "-" field is ignored

func main() {
	unmarshalExample()
	validationExample()
	complexDataExample()
	marshalExample()
	marshalWithTags()
	prettyMarshal()
}

func unmarshalExample() {
	data := []byte(`
		{
			"message": "Greetings fellow gopher",
			"otherMessage": "Good bye, little hobbit"
		}
	`)
	var v greeting
	err := json.Unmarshal(data, &v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v.Message)
	fmt.Println(v.OtherMessage)
}

func validationExample() {
	badData := []byte(`{message": "Greetings fellow gopher"}`)
	if !json.Valid(badData) {
		fmt.Printf("JSON is not valid: %s\n", badData)
	}
}

func complexDataExample() {
	complexData := []byte(`  
  	{  
			"lname": "Smith",  
			"fname": "John",  
			"address": {  
				"street": "Sulphur Springs Rd",  
				"city": "Park City",  
				"state": "VA",  
				"zipcode": 12345  
			}  
		}  
	`)
	var p person
	err := json.Unmarshal(complexData, &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", p)
}

func marshalExample() {
	var g goodbye
	g.SomeMessage = "Marshal me"
	json, err := json.Marshal(g)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", json)
}

func marshalWithTags() {
	var b book
	b.ISBN = "9933HIST"
	b.Title = "Greatest of all Books"
	b.Author = "John Adams"
	json, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", json)
}

func prettyMarshal() {
	p := person{Lastname: "Vader", Firstname: "Darth"}
	p.Address.Street = "Galaxy Far Away"
	p.Address.City = "Dark Side"
	p.Address.State = "Tatooine"
	p.Address.ZipCode = 12345
	noPrettyPrint, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	prettyPrint, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(noPrettyPrint))
	fmt.Println()
	fmt.Println(string(prettyPrint))
}