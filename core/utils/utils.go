package utils

import (
    "fmt"
	"log"
	"os"
)

var filePath = "asset/secure.gocrypted"

func ReadData() string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		if err := os.Mkdir("asset", os.ModePerm); err != nil {
			log.Fatal(err)
		}

		if err := os.WriteFile(filePath, []byte(""), 0644); err != nil {
			log.Println("Error creating file:", err)

		}
		content = []byte("")
        } 
    return string(content)
}

func AddData(newData string) {

	// Open the file in append mode, creating it if it doesn't exist
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Append the new data to the file
	_, err = file.WriteString(newData)
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}

}
