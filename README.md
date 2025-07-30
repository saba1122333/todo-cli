# Task Tracker CLI

A simple command-line task management tool written in Go. Store, track, and manage your tasks with ease.

## Features

- ✅ Add new tasks
- 📋 List all tasks or filter by status
- ✏️ Update task descriptions
- ��️ Delete tasks
- 🔄 Mark tasks as in-progress or done
- �� Persistent storage using JSON

## Installation 

### Prerequisites
- Go 1.24.5 or higher

### Build from Source 
```
git clone https://github.com/saba1122333/todo-cli
cd todo-cli 
go build -o todo-cli
```

### Run directly
``` 
go run main.go [command] [arguments]
``` 

## Usage 

# Adding a new task
``` 
go run main.go add "Buy groceries"
```


# Deleting Task 
``` 
go run main.go delete 1 
``` 

# Update Task
``` 
go run main.go update 1 "Buy groceries and cook dinner"
``` 

# Marking a task as in progress or done
``` 
go run main.go mark-in-progress 1 
go run main.go mark-done 1 
``` 

# List 
```bash 
go run main.go list
``` 

# List by status 
``` 
go run main.go list todo 
go run main.go list in-progress
go run main.go list Done
```

## Data storage 

Tasks are stored in Tasks.json file in the project root directory. the file is created automaticaly during first run of project.

## Project structure
../
    ├── cli/ # Command-line interface logic
    ├── task/ # Task management and storage
    ├── utils/ # Utility functions
    ├── main.go # Application entry point
    └── Tasks.json # Task data storage
