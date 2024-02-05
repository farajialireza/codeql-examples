package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	// Set up a simple HTTP server with a vulnerable command execution
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/execute", executeHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Display a simple HTML form for user input
	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Command Execution Demo</title>
		</head>
		<body>
			<h2>Enter a command to execute:</h2>
			<form action="/execute" method="post">
				<input type="text" name="command" required>
				<input type="submit" value="Execute">
			</form>
		</body>
		</html>
	`
	fmt.Fprintln(w, html)
}

func executeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	command := r.FormValue("command")

	// Vulnerable command execution
	output, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		fmt.Fprintf(w, "Error executing command: %v\n", err)
		return
	}

	fmt.Fprintf(w, "Command output:\n%s\n", output)
}
