package main

import (
	"strings"
)

func processCommand(command string) (string, []string) {
	if !containsNewDeclarations(command) {
		return command, nil
	}
	// new variables are declared but not used
	// call use on them to avoid "declared and not used" error

	newVariables := getNewVariables(command)
	return command, newVariables
}

func containsNewDeclarations(command string) bool {
	return strings.Contains(command, ":=")
}

func getNewVariables(command string) []string {
	variableSection := strings.Split(command, ":=")[0]
	variables := strings.Split(variableSection, ",")
	for i, variable := range variables {
		variables[i] = strings.TrimSpace(variable)
	}
	return variables
}

func isFunctionDeclaration(command string) bool {
	return strings.HasPrefix(command, "func")
}

func isVarDeclaration(command string) bool {
	return strings.HasPrefix(command, "var")
}

func isExperimentalInput(command string) bool {
	containsAssignment := strings.ContainsAny(command, ":=")
	if containsAssignment || isFunctionDeclaration(command) || isVarDeclaration(command) {
		return false
	}
	return true
}
