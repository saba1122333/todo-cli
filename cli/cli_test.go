package cli

import (
	"os"
	"testing"

	"github.com/saba1122333/todo-cli/task"
)

func destroy() {
	os.Remove(task.FileName)
	task.FileName = "Tasks.json"
}

func setup() {
	task.FileName = "cli_test.json"
	os.Remove(task.FileName)
}

func TestValidArgumentCount(t *testing.T) {

	cases := []struct {
		name      string
		args      []string
		expected  int
		expectErr bool
	}{
		{
			name:      "Correct argument count",
			args:      []string{"program", "add", "buy milk"},
			expected:  3,
			expectErr: false,
		},
		{
			name:      "Too few arguments",
			args:      []string{"program", "add"},
			expected:  3,
			expectErr: true,
		},
		{
			name:      "Too many arguments",
			args:      []string{"program", "add", "buy milk", "extra"},
			expected:  3,
			expectErr: true,
		},
		{
			name:      "Empty args",
			args:      []string{},
			expected:  3,
			expectErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Save original args
			originalArgs := os.Args
			// Restore after test
			defer func() { os.Args = originalArgs }()

			// Set test args
			os.Args = tc.args

			err := validArgumentCount(tc.expected)

			if tc.expectErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}
func TestParseID(t *testing.T) {
	cases := []struct {
		name       string
		args       []string
		argIndex   int
		expectedID int
		expectErr  bool
	}{
		{
			name:       "happy path, valid int at index",
			args:       []string{"prog", "update", "5"},
			argIndex:   2,
			expectedID: 5,
			expectErr:  false,
		},
		{
			name:       "index out of range, not enough args",
			args:       []string{"prog", "update"},
			argIndex:   2,
			expectedID: 0,
			expectErr:  true,
		},
		{
			name:       "cannot parse int, word instead of number",
			args:       []string{"prog", "update", "foo"},
			argIndex:   2,
			expectedID: -1,
			expectErr:  true,
		},
		{
			name:       "empty args",
			args:       []string{},
			argIndex:   2,
			expectedID: 0,
			expectErr:  true,
		},
		{
			name:       "empty string at index",
			args:       []string{"prog", "update", ""},
			argIndex:   2,
			expectedID: -1,
			expectErr:  true,
		},
	}

	for _, tc := range cases {
		setup()
		defer destroy()
		t.Run(tc.name, func(t *testing.T) {
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			os.Args = tc.args

			result, err := parseID(tc.argIndex)

			if tc.expectErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
			if result != tc.expectedID {
				t.Errorf("Expected ID %d, got %d", tc.expectedID, result)
			}
		})
	}
}

func TestGetDescription(t *testing.T) {
	cases := []struct {
		name         string
		args         []string
		argIndex     int
		expectedDesc string
		expectErr    bool
	}{
		{
			name:         "valid description",
			args:         []string{"prog", "add", "buy milk"},
			argIndex:     2,
			expectedDesc: "buy milk",
			expectErr:    false,
		},
		{
			name:         "description with spaces",
			args:         []string{"prog", "add", "  buy milk  "},
			argIndex:     2,
			expectedDesc: "buy milk",
			expectErr:    false,
		},
		{
			name:         "index out of range",
			args:         []string{"prog", "add"},
			argIndex:     2,
			expectedDesc: "",
			expectErr:    true,
		},
		{
			name:         "empty description",
			args:         []string{"prog", "add", ""},
			argIndex:     2,
			expectedDesc: "",
			expectErr:    true,
		},
		{
			name:         "whitespace only description",
			args:         []string{"prog", "add", "   "},
			argIndex:     2,
			expectedDesc: "",
			expectErr:    true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			os.Args = tc.args

			result, err := getDescription(tc.argIndex)

			if tc.expectErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
			if result != tc.expectedDesc {
				t.Errorf("Expected description '%s', got '%s'", tc.expectedDesc, result)
			}
		})
	}
}

func TestParseCommand(t *testing.T) {
	cases := []struct {
		name        string
		args        []string
		argIndex    int
		expectedCmd string
		expectErr   bool
	}{
		{
			name:        "valid command",
			args:        []string{"prog", "add", "buy milk"},
			argIndex:    1,
			expectedCmd: "add",
			expectErr:   false,
		},
		{
			name:        "command with spaces",
			args:        []string{"prog", "  ADD  ", "buy milk"},
			argIndex:    1,
			expectedCmd: "add",
			expectErr:   false,
		},
		{
			name:        "uppercase command",
			args:        []string{"prog", "UPDATE", "5", "new desc"},
			argIndex:    1,
			expectedCmd: "update",
			expectErr:   false,
		},
		{
			name:        "index out of range",
			args:        []string{"prog"},
			argIndex:    1,
			expectedCmd: "",
			expectErr:   true,
		},
		{
			name:        "empty args",
			args:        []string{},
			argIndex:    1,
			expectedCmd: "",
			expectErr:   true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			os.Args = tc.args

			result, err := parseCommand(tc.argIndex)

			if tc.expectErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
			if result != tc.expectedCmd {
				t.Errorf("Expected command '%s', got '%s'", tc.expectedCmd, result)
			}
		})
	}
}

func TestHandleAddCommand(t *testing.T) {
	cases := []struct {
		name      string
		args      []string
		expectErr bool
	}{
		{
			name:      "valid add command",
			args:      []string{"prog", "add", "buy milk"},
			expectErr: false,
		},
		{
			name:      "missing description",
			args:      []string{"prog", "add"},
			expectErr: true,
		},
		{
			name:      "too many arguments",
			args:      []string{"prog", "add", "buy milk", "extra"},
			expectErr: true,
		},
		{
			name:      "empty description",
			args:      []string{"prog", "add", ""},
			expectErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup()
			defer destroy()
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			os.Args = tc.args

			err := HandleAddCommand()

			if tc.expectErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func TestHandleDeleteCommand(t *testing.T) {
	cases := []struct {
		name      string
		args      []string
		expectErr bool
	}{
		{
			name:      "missing ID",
			args:      []string{"prog", "delete"},
			expectErr: true,
		},
		{
			name:      "invalid ID",
			args:      []string{"prog", "delete", "abc"},
			expectErr: true,
		},
		{
			name:      "too many arguments",
			args:      []string{"prog", "delete", "5", "extra"},
			expectErr: true,
		},
		{
			name:      "valid delete command (will fail due to empty list, but CLI validation passes)",
			args:      []string{"prog", "delete", "1"},
			expectErr: false, // CLI validation passes, but task operation will fail
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup()
			defer destroy()
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			os.Args = tc.args
			err := HandleDeleteCommand()

			if tc.expectErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.expectErr && err != nil {
				// For valid CLI args, we expect the task operation to fail due to empty list
				// but the CLI validation should pass
				if err.Error() != "to-do list is empty" && err.Error() != "task with provided id does not exist" {
					t.Errorf("Expected task operation error but got: %v", err)
				}
			}
		})
	}
}

func TestHandleUpdateCommand(t *testing.T) {
	cases := []struct {
		name      string
		args      []string
		expectErr bool
	}{
		{
			name:      "missing ID",
			args:      []string{"prog", "update", "new description"},
			expectErr: true,
		},
		{
			name:      "missing description",
			args:      []string{"prog", "update", "5"},
			expectErr: true,
		},
		{
			name:      "invalid ID",
			args:      []string{"prog", "update", "abc", "new description"},
			expectErr: true,
		},
		{
			name:      "empty description",
			args:      []string{"prog", "update", "5", ""},
			expectErr: true,
		},
		{
			name:      "too many arguments",
			args:      []string{"prog", "update", "5", "new desc", "extra"},
			expectErr: true,
		},
		{
			name:      "valid update command (will fail due to empty list, but CLI validation passes)",
			args:      []string{"prog", "update", "1", "new description"},
			expectErr: false, // CLI validation passes, but task operation will fail
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup()
			defer destroy()
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			os.Args = tc.args

			err := HandleUpdateCommand()

			if tc.expectErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.expectErr && err != nil {
				// For valid CLI args, we expect the task operation to fail due to empty list
				// but the CLI validation should pass
				if err.Error() != "to-do list is empty" && err.Error() != "task with provided id does not exist" {
					t.Errorf("Expected task operation error but got: %v", err)
				}
			}
		})
	}
}

func TestHandleListCommand(t *testing.T) {
	cases := []struct {
		name      string
		args      []string
		expectErr bool
	}{
		{
			name:      "list all",
			args:      []string{"prog", "list"},
			expectErr: false,
		},
		{
			name:      "list with status",
			args:      []string{"prog", "list", "todo"},
			expectErr: false,
		},
		{
			name:      "too many arguments",
			args:      []string{"prog", "list", "todo", "extra"},
			expectErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup()
			defer destroy()
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			os.Args = tc.args

			err := HandleListCommand()

			if tc.expectErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func TestHandleMarkCommand(t *testing.T) {
	cases := []struct {
		name      string
		command   string
		args      []string
		expectErr bool
	}{
		{
			name:      "invalid mark command",
			command:   "mark-invalid",
			args:      []string{"prog", "mark-invalid", "1"},
			expectErr: true,
		},
		{
			name:      "missing ID",
			command:   "mark-in-progress",
			args:      []string{"prog", "mark-in-progress"},
			expectErr: true,
		},
		{
			name:      "invalid ID",
			command:   "mark-in-progress",
			args:      []string{"prog", "mark-in-progress", "abc"},
			expectErr: true,
		},
		{
			name:      "too many arguments",
			command:   "mark-in-progress",
			args:      []string{"prog", "mark-in-progress", "5", "extra"},
			expectErr: true,
		},
		{
			name:      "valid mark in-progress command (will fail due to empty list, but CLI validation passes)",
			command:   "mark-in-progress",
			args:      []string{"prog", "mark-in-progress", "1"},
			expectErr: false, // CLI validation passes, but task operation will fail
		},
		{
			name:      "valid mark done command (will fail due to empty list, but CLI validation passes)",
			command:   "mark-done",
			args:      []string{"prog", "mark-done", "1"},
			expectErr: false, // CLI validation passes, but task operation will fail
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup()
			defer destroy()
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			os.Args = tc.args

			err := HandleMarkCommand(tc.command)

			if tc.expectErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.expectErr && err != nil {
				// For valid CLI args, we expect the task operation to fail due to empty list
				// but the CLI validation should pass
				if err.Error() != "to-do list is empty" && err.Error() != "task with provided id does not exist" {
					t.Errorf("Expected task operation error but got: %v", err)
				}
			}
		})
	}
}

func TestRun(t *testing.T) {
	cases := []struct {
		name      string
		args      []string
		expectErr bool
	}{
		{
			name:      "valid add command",
			args:      []string{"prog", "add", "buy milk"},
			expectErr: false,
		},
		{
			name:      "valid delete command",
			args:      []string{"prog", "delete", "1"},
			expectErr: false,
		},
		{
			name:      "valid update command",
			args:      []string{"prog", "update", "1", "new description"},
			expectErr: false,
		},
		{
			name:      "valid list command",
			args:      []string{"prog", "list"},
			expectErr: false,
		},
		{
			name:      "valid mark-in-progress command",
			args:      []string{"prog", "mark-in-progress", "1"},
			expectErr: false,
		},
		{
			name:      "valid mark-done command",
			args:      []string{"prog", "mark-done", "1"},
			expectErr: false,
		},
		{
			name:      "missing command",
			args:      []string{"prog"},
			expectErr: true,
		},
		{
			name:      "unknown command",
			args:      []string{"prog", "unknown"},
			expectErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setup()
			defer destroy()
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			os.Args = tc.args

			err := Run()

			if tc.expectErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.expectErr && err != nil {
				// For valid CLI args, we expect the task operation to fail due to empty list
				// but the CLI validation should pass
				if err.Error() != "to-do list is empty" && err.Error() != "task with provided id does not exist" {
					t.Errorf("Expected task operation error but got: %v", err)
				}
			}
		})
	}
}
