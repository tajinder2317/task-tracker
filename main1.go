// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	"time"
// )

// type Task struct {
// 	ID          int       `json:"id"`
// 	Description string    `json:"description"`
// 	Status      string    `json:"status"` // "todo", "in-progress", "done"
// 	CreatedAt   time.Time `json:"createdAt"`
// 	UpdatedAt   time.Time `json:"updatedAt"`
// }

// const dbFileName = "tasks.json"

// func main() {
// 	if _, err := os.Stat(dbFileName); os.IsNotExist(err) {
// 		err := os.WriteFile(dbFileName, []byte("[]"), 0644)
// 		if err != nil {
// 			fmt.Printf("Error creating the storage file %v\n", err)
// 			os.Exit(1)
// 		}
// 	}

// 	if len(os.Args) < 2 {
// 		printUsage()
// 		return
// 	}
	
// 	command := os.Args[1]

// 	// 3. Route arguments to their respected features
// 	switch command {
// 	case "add":
// 		if len(os.Args) < 3 {
// 			fmt.Println("Error: Missing Task description.  Usage: task-cli add \"description\"")
// 			return
// 		}
// 		addTask(os.Args[2])
// 	case "list":
// 		status := ""
// 		if len(os.Args) == 3 {
// 			status = os.Args[2] // e.g., "done", "todo", "in-progress"
// 		}
// 		listTasks(status)

// 		// TODO: Implement cases for "update", "delete", "mark-in-progress", "mark-done"

// 	default:
// 		fmt.Printf("Unknown command: %s\n",command)
// 		printUsage()
// 	}

// }

// func printUsage(){
// 	fmt.Println("Usage task-cli <command> [arguments]")
// 	fmt.Println("Commands:")
// 	fmt.Println(" list [status]")
// }


// func loadTask() {
// 	data,err := os.ReadFile(dbFileName)
// 	if _, err != nil{
// 		return nil, err
// 	}
// 	var task []Task
// 	err = json.Unmarshal(data,&tasks)
// 	return tasks, err
// }
// func updateTask() {

// }
// func deleteTask() {}
