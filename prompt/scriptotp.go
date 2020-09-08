package prompt

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// ScriptMfaProvider executes a user provided script that returns a opt token
func ScriptMfaProvider(mfaSerial string) (string, error) {
	scriptName := os.Getenv("AWS_VAULT_MFA_SCRIPT")
	if scriptName == "" {
		return "", fmt.Errorf("AWS_VAULT_MFA_SCRIPT not defined in environment")
	}

	log.Printf("Fetching MFA code using `%s`", scriptName)
	cmd := exec.Command(scriptName)
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("scriptmfa: %w", err)
	}

	return strings.TrimSpace(string(out)), nil
}

func init() {
	Methods["scriptmfa"] = ScriptMfaProvider
}
