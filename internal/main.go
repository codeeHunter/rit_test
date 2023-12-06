package internal

import (
	"fmt"
	"time"

	"github.com/codeeHunter/rit_test/utils"
)

type Action struct {
	Type       string                 `json:"type"`
	Name       string                 `json:"name"`
	Result     string                 `json:"result"`
	Parameters map[string]interface{} `json:"parameters"`
	Yes        []Action               `json:"yes,omitempty"`
	No         []Action               `json:"no,omitempty"`
}

type Sequence struct {
	Sequence []Action `json:"sequence"`
}

func executeAction(action *Action) {
	switch action.Type {
	case "action":
		switch action.Name {
		case "CreateFile":
			filename := action.Parameters["filename"].(string)
			err := utils.CreateFile(filename)
			if err != nil {
				action.Result = fmt.Sprintf("Error creating file: %s", err)
			} else {
				action.Result = "File created successfully"
			}
		case "ChangeFileName":
			filename := action.Parameters["filename"].(string)
			newFilename := action.Parameters["new_filename"].(string)
			err := utils.ChangeFileName(filename, newFilename)
			if err != nil {
				action.Result = fmt.Sprintf("Error changing filename: %s", err)
			} else {
				action.Result = "Filename changed successfully"
			}
		case "DeleteFile":
			filename := action.Parameters["filename"].(string)
			err := utils.DeleteFile(filename)
			if err != nil {
				action.Result = fmt.Sprintf("Error deleting file: %s", err)
			} else {
				action.Result = "File deleted successfully"
			}
		case "GetFileCreationTime":
			filename := action.Parameters["filename"].(string)
			creationTime, err := utils.GetFileCreationTime(filename)
			if err != nil {
				action.Result = fmt.Sprintf("Error getting file creation time: %s", err)
			} else {
				action.Result = fmt.Sprintf("File creation time: %s", creationTime)
			}
		case "WriteToFile":
			filename := action.Parameters["filename"].(string)
			content := action.Parameters["content"].(string)
			err := utils.WriteToFile(filename, content)
			if err != nil {
				action.Result = fmt.Sprintf("Error writing to file: %s", err)
			} else {
				action.Result = "Content written to file successfully"
			}
		}
	}
}

func ExecuteSequence(sequence []Action) {
	for i := range sequence {
		executeAction(&sequence[i])
		if sequence[i].Type == "condition" {
			conditionValue := sequence[i].Parameters["value"].(string)
			conditionTime, err := time.Parse(time.RFC3339, conditionValue)
			if err != nil {
				fmt.Println("Error parsing condition time:", err)
			}
			currentTime := time.Now()
			if currentTime.After(conditionTime) {
				ExecuteSequence(sequence[i].Yes)
			} else {
				ExecuteSequence(sequence[i].No)
			}
		}
	}
}
