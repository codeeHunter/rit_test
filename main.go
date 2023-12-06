package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/codeeHunter/rit_test/internal"
)

func main() {
	// Read JSON file
	file, err := os.ReadFile("example.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Decode JSON
	var sequence internal.Sequence
	err = json.Unmarshal(file, &sequence)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	internal.ExecuteSequence(sequence.Sequence)

	updatedJSON, err := json.MarshalIndent(sequence, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	err = os.WriteFile("updated_example.json", updatedJSON, 0644)
	if err != nil {
		fmt.Println("Error writing updated JSON file:", err)
		return
	}

	fmt.Println("Execution completed successfully.")
}
