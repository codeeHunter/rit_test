package utils

import (
	"fmt"
	"os"
	"time"
)

func CreateFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()
	return nil
}

func ChangeFileName(filename, newFilename string) error {
	err := os.Rename(filename, newFilename)
	if err != nil {
		fmt.Println("Error changing filename:", err)
		return err
	}
	return nil
}

func DeleteFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return err
	}
	return nil
}

func GetFileCreationTime(filename string) (string, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println("Error getting file creation time:", err)
		return "", err
	}
	creationTime := fileInfo.ModTime().Format(time.RFC3339)
	return creationTime, nil
}

func WriteToFile(filename, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	return nil
}
