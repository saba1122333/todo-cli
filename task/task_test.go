package task

import (
	"os"
	"testing"
)

func destroy() {
	os.Remove(FileName)
	FileName = "Tasks.json"
}

func setup() {
	FileName = "tasks_test.json"
	os.Remove(FileName)
}

func TestReadWriteTasks(t *testing.T) {

	cases := []struct {
		name  string
		tasks []Task
	}{
		{
			name:  "Empty",
			tasks: []Task{},
		},
		{
			name: "Single",
			tasks: []Task{
				{ID: 1, Description: "buy milk", Status: Todo, CreatedAt: "", UpdatedAt: " "},
			},
		},
		{
			name: "Multiple",
			tasks: []Task{
				{ID: 2, Description: "buy milk", Status: Todo, CreatedAt: "", UpdatedAt: " "},
				{ID: 3, Description: "buy ball", Status: Todo, CreatedAt: "", UpdatedAt: " "},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup()
			defer destroy()
			err := WriteTasks(tc.tasks)
			if err != nil {
				t.Fatalf("Failed to Write %v", err)
			}
			written, err := ReadTasks()
			if err != nil {
				t.Fatalf("Failed to Read %v", err)
			}

			if len(written) != len(tc.tasks) {
				t.Errorf("Expected %d tasks, got %d", len(tc.tasks), len(written))
			}
			for i, expected := range tc.tasks {
				if i > len(written) {
					t.Errorf("Missing task at index %d", i)
					continue
				}
				if expected != written[i] {
					t.Errorf("Expected %v does not match Read value %v", tc.tasks[i], expected)
				}
			}

		})
	}
}

func TestAppendTask(t *testing.T) {
	setup()
	defer destroy()
	cases := []struct {
		name        string
		description string
	}{
		{
			name:        "Empty File Append",
			description: "buy milk",
		},
		{
			name:        "Non-Empty File Append",
			description: "buy milk",
		},
	}
	err := AppendTask(cases[0].description)
	if err != nil {
		t.Errorf("Failed to Append %v ", err)
	}
	written, err := ReadTasks()
	if err != nil {
		t.Errorf("Failed to Read %v ", err)
	}
	if written[0].Description != cases[0].description {
		t.Fatalf("Failed on case: %v Expected %v, got %v", cases[0].name, cases[0].description, written[0].Description)
	}

	err = AppendTask(cases[1].description)
	if err != nil {
		t.Fatalf("Failed to Append %v ", err)
	}
	written, err = ReadTasks()
	if err != nil {
		t.Errorf("Failed to Read %v ", err)
	}
	if written[1].Description != cases[1].description {
		t.Fatalf("Failed on case: %v Expected %v, got %v", cases[1].name, cases[1].description, written[1].Description)
	}

}
func TestDeleteTask(t *testing.T) {
	cases := []struct {
		name            string
		setupTasks      []string
		deleteID        int
		expectError     bool
		expectRemaining int
	}{
		{
			name:            "Delete from multiple tasks",
			setupTasks:      []string{"buy milk", "buy bread", "buy eggs"},
			deleteID:        2,
			expectError:     false,
			expectRemaining: 2,
		},
		{
			name:            "Delete from single task",
			setupTasks:      []string{"buy milk"},
			deleteID:        1,
			expectError:     false,
			expectRemaining: 0,
		},
		{
			name:            "Delete non-existent ID",
			setupTasks:      []string{"buy milk"},
			deleteID:        999,
			expectError:     true,
			expectRemaining: 1,
		},
		{
			name:            "Delete from empty list",
			setupTasks:      []string{},
			deleteID:        1,
			expectError:     true,
			expectRemaining: 0,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup()
			defer destroy()
			for _, dc := range tc.setupTasks {
				err := AppendTask(dc)
				if err != nil {
					t.Fatalf("Failed to setupTasks: %v", err)
				}
			}
			err := DeleteTask(tc.deleteID)

			if tc.expectError && err == nil {
				t.Fatalf("Expected error but got none: %v", err)
			}
			if !tc.expectError && err != nil {
				t.Fatalf("Did not Expected error but got one: %v", err)
			}
			remaining, err := ReadTasks()

			if err != nil {
				t.Fatalf("Failed to Read: %v", err)
			}
			if len(remaining) != tc.expectRemaining {
				t.Fatalf("Expect Remaining %v remaining tasks, got %v ", tc.expectRemaining, len(remaining))
			}
			for _, ts := range remaining {
				if ts.ID == tc.deleteID {
					t.Fatalf("Deleted Task is not deleted: %v", err)

				}
			}

		})
	}
}

func TestUpdateTask(t *testing.T) {
	cases := []struct {
		name        string
		setupTasks  []string
		updateID    int
		description string
		expectError bool
	}{
		{
			name:        "Update from multiple tasks",
			setupTasks:  []string{"buy milk", "buy bread", "buy eggs"},
			updateID:    2,
			description: "buy bread-2 ",
			expectError: false,
		},
		{
			name:        "Update from single task",
			setupTasks:  []string{"buy milk"},
			updateID:    1,
			description: "buy milk-2 ",
			expectError: false,
		},
		{
			name:        "Update non-existent ID",
			setupTasks:  []string{"buy milk"},
			updateID:    999,
			description: "buy milk-2 ",
			expectError: true,
		},
		{
			name:        "Update from empty list",
			setupTasks:  []string{},
			updateID:    1,
			expectError: true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup()
			defer destroy()
			for _, dc := range tc.setupTasks {
				err := AppendTask(dc)
				if err != nil {
					t.Fatalf("Failed to setupTasks: %v", err)
				}
			}
			err := UpdateTask(tc.updateID, tc.description)

			if tc.expectError && err == nil {
				t.Fatalf("Expected error but got none: %v", err)
			}
			if !tc.expectError && err != nil {
				t.Fatalf("Did not Expected error but got one: %v", err)
			}
			remaining, err := ReadTasks()

			if err != nil {
				t.Fatalf("Failed to Read: %v", err)
			}
			if len(remaining) != len(tc.setupTasks) {
				t.Fatalf("Expect Remaining %v remaining tasks, got %v ", len(tc.setupTasks), len(remaining))
			}
			for _, ts := range remaining {

				if ts.ID == tc.updateID && ts.Description != tc.description {
					t.Fatalf("Updated Task is not Updated: %v", ts.ID)
				}
			}

		})
	}
}

func TestMarkTask(t *testing.T) {
	cases := []struct {
		name        string
		setupTasks  []string
		MarkID      int
		status      Status
		expectError bool
	}{
		{
			name:        "Mark single task as todo",
			setupTasks:  []string{"buy milk"},
			MarkID:      1,
			status:      Todo,
			expectError: false,
		},
		{
			name:        "Mark single task as in-progress",
			setupTasks:  []string{"buy milk"},
			MarkID:      1,
			status:      Inprogress,
			expectError: false,
		},
		{
			name:        "Mark single task as done",
			setupTasks:  []string{"buy milk"},
			MarkID:      1,
			status:      Done,
			expectError: false,
		},
		{
			name:        "Mark one of multiple tasks as todo",
			setupTasks:  []string{"buy milk", "buy bread", "buy eggs"},
			MarkID:      2,
			status:      Todo,
			expectError: false,
		},
		{
			name:        "Mark one of multiple tasks as in-progress",
			setupTasks:  []string{"buy milk", "buy bread", "buy eggs"},
			MarkID:      2,
			status:      Inprogress,
			expectError: false,
		},
		{
			name:        "Mark one of multiple tasks as done",
			setupTasks:  []string{"buy milk", "buy bread", "buy eggs"},
			MarkID:      2,
			status:      Done,
			expectError: false,
		},
		{
			name:        "Mark non-existent ID",
			setupTasks:  []string{"buy milk"},
			MarkID:      999,
			status:      Inprogress,
			expectError: true,
		},
		{
			name:        "Mark in empty list",
			setupTasks:  []string{},
			MarkID:      999,
			status:      Inprogress,
			expectError: true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup()
			defer destroy()
			for _, dc := range tc.setupTasks {
				if err := AppendTask(dc); err != nil {
					t.Fatalf("Failed to setupTasks: %v", err)
				}
			}
			err := MarkTask(tc.MarkID, tc.status)

			if tc.expectError && err == nil {
				t.Fatalf("Expected error but got none")
			}
			if !tc.expectError && err != nil {
				t.Fatalf("Did not expect error but got one: %v", err)
			}
			remaining, err := ReadTasks()
			if err != nil {
				t.Fatalf("Failed to Read: %v", err)
			}
			if len(remaining) != len(tc.setupTasks) {
				t.Fatalf("Expected %v remaining tasks, got %v", len(tc.setupTasks), len(remaining))
			}
			for _, ts := range remaining {
				if ts.ID == tc.MarkID && !tc.expectError && ts.Status != tc.status {
					t.Fatalf("Task %v not marked correctly: got %v, want %v", ts.ID, ts.Status, tc.status)
				}
				if ts.ID == tc.MarkID && tc.expectError && ts.Status == tc.status {
					t.Fatalf("Task %v should not have been marked", ts.ID)
				}
			}
		})
	}
}
func TestListTask(t *testing.T) {

	cases := []struct {
		name          string
		setupTasks    []Task
		status        string
		ExpectedTasks []Task
		ExpectedErr   bool
	}{
		{
			name:          "List empty list",
			setupTasks:    []Task{},
			status:        "todo",
			ExpectedTasks: []Task{},
			ExpectedErr:   false,
		},

		{
			name: "List Single Task ",
			setupTasks: []Task{
				{ID: 1, Description: "buy milk", Status: Todo, CreatedAt: "", UpdatedAt: " "},
			},
			status: "todo",
			ExpectedTasks: []Task{
				{ID: 1, Description: "buy milk", Status: Todo, CreatedAt: "", UpdatedAt: " "},
			},
			ExpectedErr: false,
		},
		{
			name: "List Single Task With wrong status",
			setupTasks: []Task{
				{ID: 1, Description: "buy milk", Status: Todo, CreatedAt: "", UpdatedAt: " "},
			},
			status:        "in-progress",
			ExpectedTasks: []Task{},
			ExpectedErr:   false,
		},

		{
			name: "List Single Task With invalid status",
			setupTasks: []Task{
				{ID: 1, Description: "buy milk", Status: Todo, CreatedAt: "", UpdatedAt: " "},
			},
			status:        "ready",  // This is not a valid status. Should trigger an error.
			ExpectedTasks: []Task{}, // No tasks should be returned.
			ExpectedErr:   true,     // Error is expected.
		},
		{
			name: "List Single Task With Inprogress Status",
			setupTasks: []Task{
				{ID: 1, Description: "buy milk", Status: Inprogress, CreatedAt: "", UpdatedAt: " "},
			},
			status: "in-progress",
			ExpectedTasks: []Task{
				{ID: 1, Description: "buy milk", Status: Inprogress, CreatedAt: "", UpdatedAt: " "},
			},
			ExpectedErr: false,
		},
		{
			name: "List Single Task With all Status",
			setupTasks: []Task{
				{ID: 1, Description: "buy milk", Status: Done, CreatedAt: "", UpdatedAt: " "},
			},
			status: "all",
			ExpectedTasks: []Task{
				{ID: 1, Description: "buy milk", Status: Done, CreatedAt: "", UpdatedAt: " "},
			},
			ExpectedErr: false,
		},
		{
			name: "List multiple Task With all Status",
			setupTasks: []Task{
				{ID: 1, Description: "buy milk", Status: Todo, CreatedAt: "", UpdatedAt: " "},
				{ID: 2, Description: "buy milk", Status: Inprogress, CreatedAt: "", UpdatedAt: " "},
				{ID: 3, Description: "buy milk", Status: Done, CreatedAt: "", UpdatedAt: " "},
			},
			status: "all",
			ExpectedTasks: []Task{
				{ID: 1, Description: "buy milk", Status: Todo, CreatedAt: "", UpdatedAt: " "},
				{ID: 2, Description: "buy milk", Status: Inprogress, CreatedAt: "", UpdatedAt: " "},
				{ID: 3, Description: "buy milk", Status: Done, CreatedAt: "", UpdatedAt: " "},
			},
			ExpectedErr: false,
		},
		{
			name: "List multiple Task With Todo Status",
			setupTasks: []Task{
				{ID: 1, Description: "buy milk", Status: Todo, CreatedAt: "", UpdatedAt: " "},
				{ID: 2, Description: "buy milk", Status: Inprogress, CreatedAt: "", UpdatedAt: " "},
				{ID: 3, Description: "buy milk", Status: Done, CreatedAt: "", UpdatedAt: " "},
			},
			status: "todo",
			ExpectedTasks: []Task{
				{ID: 1, Description: "buy milk", Status: Todo, CreatedAt: "", UpdatedAt: " "},
			},
			ExpectedErr: false,
		},
		{
			name: "List multiple Task With Done Status",
			setupTasks: []Task{
				{ID: 1, Description: "buy milk", Status: Todo, CreatedAt: "", UpdatedAt: " "},
				{ID: 2, Description: "buy milk", Status: Inprogress, CreatedAt: "", UpdatedAt: " "},
				{ID: 3, Description: "buy milk", Status: Done, CreatedAt: "", UpdatedAt: " "},
			},
			status: "done",
			ExpectedTasks: []Task{
				{ID: 3, Description: "buy milk", Status: Done, CreatedAt: "", UpdatedAt: " "},
			},
			ExpectedErr: false,
		},
		{
			name: "List multiple Task With Inprogress Status",
			setupTasks: []Task{
				{ID: 1, Description: "buy milk", Status: Todo, CreatedAt: "", UpdatedAt: " "},
				{ID: 2, Description: "buy milk", Status: Inprogress, CreatedAt: "", UpdatedAt: " "},
				{ID: 3, Description: "buy milk", Status: Done, CreatedAt: "", UpdatedAt: " "},
			},
			status: "in-progress",
			ExpectedTasks: []Task{
				{ID: 2, Description: "buy milk", Status: Inprogress, CreatedAt: "", UpdatedAt: " "},
			},
			ExpectedErr: false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup()
			defer destroy()
			if err := WriteTasks(tc.setupTasks); err != nil {
				t.Fatalf("Failed to setUpTasks %v", err)
			}
			filtered, err := ListTasks(tc.status)

			if tc.ExpectedErr && err == nil {
				t.Fatalf("Expected error but got none")
			}
			if !tc.ExpectedErr && err != nil {
				t.Fatalf("Did not expect error but got one: %v", err)
			}
			if len(filtered) != len(tc.ExpectedTasks) {
				t.Fatalf("Expected %v filteredTasks tasks, got %v", len(tc.ExpectedTasks), len(filtered))
			}
			for i := range tc.ExpectedTasks {
				if filtered[i] != tc.ExpectedTasks[i] {
					t.Errorf("Expected %v, got %v", tc.ExpectedTasks[i], filtered[i])
				}
			}

		})
	}

}
func TestGenerateID(t *testing.T) {
	cases := []struct {
		name       string
		setUpTasks []Task
		expected   int
	}{
		{"Empty list", []Task{}, 1},
		{"Single task", []Task{{ID: 1}}, 2},
		{"Multiple tasks", []Task{{ID: 1}, {ID: 5}, {ID: 3}}, 6},
		{"Non-sequential", []Task{{ID: 10}, {ID: 2}, {ID: 7}}, 11},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup()
			defer destroy()
			if err := WriteTasks(tc.setUpTasks); err != nil {
				t.Fatalf("Failed to setUpTasks %v", err)
			}

			if ID := generateID(); tc.expected != ID {
				t.Errorf("Expected %d, got %d", tc.expected, ID)

			}
		})
	}
}
