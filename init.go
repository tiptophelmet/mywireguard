package main

import (
	"log"
	"os/exec"
)

func validateDependencies() {
	// Check for required dependencies
	dependencies := []string{"wg", "terraform"}

	for _, dep := range dependencies {
		if _, err := exec.LookPath(dep); err != nil {
			log.Fatalf("'%s' executable not found in PATH. Please install the required dependency", dep)
		}
	}
}
