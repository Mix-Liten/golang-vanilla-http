package utils

import (
	"fmt"
	"io/ioutil"
)

type Note struct {
	Title string
	Body  []byte
}

func getInput(field string) string {
	var input string
	fmt.Println("Enter file " + field + ":")
	fmt.Scan(&input)
	return input
}

func CreateNote() *Note {
	title := getInput("title")
	bodyTxt := getInput("body")
	return &Note{
		Title: title,
		Body:  []byte(bodyTxt),
	}
}

func (p *Note) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func LoadNote(title string) (*Note, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Note{Title: title, Body: body}, nil
}

//func main() {
//	p1 := utils.CreateNote()
//	p1.Save()
//	p2, _ := utils.LoadNote(p1.Title)
//	fmt.Println(string(p2.Body))
//}
