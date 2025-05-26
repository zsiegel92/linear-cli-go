package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

// CopyToClipboard copies text to system clipboard
func CopyToClipboard(text string) error {
	var cmd *exec.Cmd
	
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "linux":
		cmd = exec.Command("xclip", "-selection", "clipboard")
	case "windows":
		cmd = exec.Command("cmd", "/c", "clip")
	default:
		return fmt.Errorf("clipboard not supported on %s", runtime.GOOS)
	}

	// Get stdin pipe before starting the process
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdin pipe: %w", err)
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		stdin.Close()
		return fmt.Errorf("failed to start clipboard command: %w", err)
	}

	// Write the text and close stdin
	_, err = stdin.Write([]byte(text))
	stdin.Close()
	if err != nil {
		return fmt.Errorf("failed to write to clipboard: %w", err)
	}

	// Wait for the command to complete
	return cmd.Wait()
}

// OpenInBrowser opens URL in default browser
func OpenInBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		return fmt.Errorf("browser opening not supported on %s", runtime.GOOS)
	}

	return cmd.Start()
}