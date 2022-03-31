package localfile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
)

// GetJSON reads the file at filePath, and Unmarshals it into the provided data structure ds.
func GetJSON(filePath string, ds interface{}) error {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Error loading secrets json: %v \n", err)
	}

	fmt.Println("printing filebytes")
	fmt.Println(string(fileBytes[:]))

	err = json.Unmarshal(fileBytes, &ds)
	if err != nil {
		return fmt.Errorf("Error Unmarshaling fileBytes to json: %v \n", err)
	}

	return nil
}

// GetString reads the file at filePath, returns a string of its contents.
func GetString(filePath string) (string, error) {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("Error reading file string: %v \n", err)
	}

	fileString := string(fileBytes[:])
	fmt.Println("printing fileString")
	fmt.Println(fileString)

	return fileString, nil
}

func GetTemplate(path string, templateOptions map[string]string) (string, error) {
	buf := new(bytes.Buffer)
	TextBody, err := template.ParseFiles(path)
	if err != nil {
		return "", fmt.Errorf("error parsing activation template file: %v\n", err)
	}

	err = TextBody.Execute(buf, templateOptions)
	if err != nil {
		return "", fmt.Errorf("error executing email template file: %v\n", err)
	}
	return buf.String(), nil
}
