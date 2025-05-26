package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/zach/linear_cli_go/cmd"
	"github.com/zach/linear_cli_go/config"
)

func main() {
	// Set up signal handling for graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\n👋 Goodbye!")
		os.Exit(0)
	}()

	// Parse command line flags - both long and short forms
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

	// Load environment variables from .env file
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	// Check if LINEAR_API_KEY is set when not using mock data
	if !useMock && os.Getenv("LINEAR_API_KEY") == "" {
		log.Fatalf("LINEAR_API_KEY environment variable is required. Please set it in .env file or environment.")
	}

	// Run the issue selector
	cmd.RunIssueSelector(useMock, onlyMine, projectID)
}
