package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// 1. Move the struct OUTSIDE of main so all functions can see it
type Task struct {
	Description string `json:"description"`
	Status      string `json:"status"`
	ID          int    `json:"id"`
}

func main() {
	// Safety check: ensure the user typed at least one command
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		// Safety check: make sure they provided a description string
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing Task description. Usage: task-cli add \"description\"")
			return
		}
		description := os.Args[2]
		addTask(description)

	case "list":
		listTasks()

	case "delete":
		// Safety check: make sure they provided an ID to delete
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing Task ID. Usage: task-cli delete <id>")
			return
		}
		id := os.Args[2]
		deleteTask(id)

	case "update":
		// Safety check: make sure they provided an ID and a new description text
		if len(os.Args) < 4 {
			fmt.Println("Error: Missing arguments. Usage: task-cli update <id> \"new description\"")
			return
		}
		id := os.Args[2]
		newDescription := os.Args[3]
		updateTask(id, newDescription)

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing Task ID. Usage: task-cli mark-in-progress <id>")
			return
		}
		id := os.Args[2]
		changeTaskStatus(id, "in-progress")

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing Task ID. Usage: task-cli mark-done <id>")
			return
		}
		id := os.Args[2]
		changeTaskStatus(id, "done")

	default:
		fmt.Println("Unknown command!")
	}
}

// 2. This function now matches the empty call from main()

func addTask(text string) {
	var taskList []Task // This holds our collection of multiple tasks

	// 1. Read existing file if it exists
	fileData, err := os.ReadFile("tasks.json")
	if err == nil {
		// File exists! Unpack the existing JSON text back into our taskList slice
		err = json.Unmarshal(fileData, &taskList)
		if err != nil {
			fmt.Println("Error reading existing tasks:", err)
			return
		}
	}

	// 2. Calculate a smart auto-incrementing ID
	newID := 1
	if len(taskList) > 0 {
		// Look at the very last task in the list and add 1 to its ID
		newID = taskList[len(taskList)-1].ID + 1
	}

	// 3. Construct your new task
	newTask := Task{
		ID:          newID,
		Description: text,
		Status:      "todo",
	}

	// 4. Append the new task to your existing list collection
	taskList = append(taskList, newTask)

	// 5. Convert the ENTIRE updated list back to JSON text
	jsonData, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		fmt.Println("Error converting list to JSON:", err)
		return
	}

	// 6. Save the entire updated collection back to tasks.json
	err = os.WriteFile("tasks.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("Task added successfully (ID: %d)\n", newID)
}

// 3. Removed 'text string' since list doesn't strictly require it on startup
func listTasks() {
	// 1. Check what filter status the user typed (if any)
	statusFilter := ""
	if len(os.Args) == 3 {
		statusFilter = os.Args[2] // e.g., "done", "todo", "in-progress"
	}

	// 2. Read the tasks.json file
	fileData, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("No tasks found. Try adding one first!")
		return
	}

	// 3. Unpack the JSON bytes into a Go slice
	var taskList []Task
	err = json.Unmarshal(fileData, &taskList)
	if err != nil {
		fmt.Println("Error reading tasks data:", err)
		return
	}

	if len(taskList) == 0 {
		fmt.Println("Your to-do list is empty.")
		return
	}

	// 4. Loop through each task and print it
	fmt.Println("-------------------------------------------")
	for _, task := range taskList {
		// If a filter was requested, skip tasks that don't match it
		if statusFilter != "" && task.Status != statusFilter {
			continue
		}

		// Print the task in a clean format
		fmt.Printf("ID: %d | [%s] | %s\n", task.ID, task.Status, task.Description)
	}
	fmt.Println("-------------------------------------------")
}

// 4. Renamed from 'delete' to 'deleteTask' and added the 'string' data type

func deleteTask(idStr string) {
	// 1. Convert the ID text input into a real integer number
	targetID, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Error: Task ID must be a valid number.")
		return
	}

	// 2. Read the existing JSON file
	fileData, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("Error: No tasks file found.")
		return
	}

	var taskList []Task
	err = json.Unmarshal(fileData, &taskList)
	if err != nil {
		fmt.Println("Error reading tasks data:", err)
		return
	}

	// 3. Loop to find the index of the task we want to delete
	foundIndex := -1
	for index, task := range taskList {
		if task.ID == targetID {
			foundIndex = index
			break
		}
	}

	// If we went through the whole loop and didn't find the ID
	if foundIndex == -1 {
		fmt.Printf("Error: Task with ID %d not found.\n", targetID)
		return
	}

	// 4. Remove the item from the slice using Go's cut-and-join trick
	taskList = append(taskList[:foundIndex], taskList[foundIndex+1:]...)

	// 5. Save the modified list back to the file
	jsonData, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		fmt.Println("Error updating tasks file:", err)
		return
	}

	err = os.WriteFile("tasks.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing changes to file:", err)
		return
	}

	fmt.Printf("Task ID %d deleted successfully.\n", targetID)
}
func updateTask(idStr string, newDesc string) {
	targetID, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Error: Task ID must be a valid number.")
		return
	}

	fileData, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("Error: No tasks file found.")
		return
	}

	var taskList []Task
	json.Unmarshal(fileData, &taskList)

	foundIndex := -1
	for index, task := range taskList {
		if task.ID == targetID {
			foundIndex = index
			break
		}
	}

	if foundIndex == -1 {
		fmt.Printf("Error: Task with ID %d not found.\n", targetID)
		return
	}

	// Overwrite the description directly at that index location
	taskList[foundIndex].Description = newDesc

	// Save back to file
	jsonData, _ := json.MarshalIndent(taskList, "", "  ")
	os.WriteFile("tasks.json", jsonData, 0644)

	fmt.Printf("Task ID %d updated successfully.\n", targetID)
}

func changeTaskStatus(idStr string, newStatus string) {
	targetID, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Error: Task ID must be a valid number.")
		return
	}

	fileData, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("Error: No tasks file found.")
		return
	}

	var taskList []Task
	json.Unmarshal(fileData, &taskList)

	foundIndex := -1
	for index, task := range taskList {
		if task.ID == targetID {
			foundIndex = index
			break
		}
	}

	if foundIndex == -1 {
		fmt.Printf("Error: Task with ID %d not found.\n", targetID)
		return
	}

	// Overwrite the status directly at that index location
	taskList[foundIndex].Status = newStatus

	// Save back to file
	jsonData, _ := json.MarshalIndent(taskList, "", "  ")
	os.WriteFile("tasks.json", jsonData, 0644)

	fmt.Printf("Task ID %d marked as %s successfully.\n", targetID, newStatus)
}
