package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/zach/linear_cli_go/cmd"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\n👋 Goodbye!")
		os.Exit(0)
	}()

	var useMock bool
	flag.BoolVar(&useMock, "d", false, "Use demo data instead of Linear API")
	flag.BoolVar(&useMock, "demo", false, "Use demo data instead of Linear API")
	
	var onlyMine bool
	flag.BoolVar(&onlyMine, "m", false, "Show only issues assigned to me")
	flag.BoolVar(&onlyMine, "mine", false, "Show only issues assigned to me")
	
	var projectID string
	flag.StringVar(&projectID, "p", "", "Filter issues by project ID")
	flag.StringVar(&projectID, "project", "", "Filter issues by project ID")
	
	var showHelp bool
	flag.BoolVar(&showHelp, "h", false, "Show help")
	flag.BoolVar(&showHelp, "help", false, "Show help")
	
	flag.Parse()

	if showHelp {
		fmt.Println("Linear CLI - A command-line interface for Linear")
		fmt.Println("\nUsage:")
		fmt.Println("  linear_cli [flags]")
		fmt.Println("\nFlags:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if !useMock && os.Getenv("LINEAR_API_KEY") == "" {
		fmt.Println("❌ LINEAR_API_KEY environment variable is required.")
		fmt.Println("")
		fmt.Println("To set it up:")
		fmt.Println("1. Get your Linear API key from https://linear.app/settings/api")
		fmt.Println("2. Add it to your shell configuration:")
		fmt.Println("   echo 'export LINEAR_API_KEY=\"your_api_key_here\"' >> ~/.zshrc")
		fmt.Println("   source ~/.zshrc")
		fmt.Println("")
		fmt.Println("Or set it temporarily:")
		fmt.Println("   export LINEAR_API_KEY=\"your_api_key_here\"")
		os.Exit(1)
	}

	cmd.RunIssueSelector(useMock, onlyMine, projectID)
}
