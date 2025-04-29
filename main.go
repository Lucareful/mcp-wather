package main

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"
	_ "go.uber.org/automaxprocs"

	"mcp-weather/internal/city"
	"mcp-weather/pkg/mcp"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"Weather Demo",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
		server.WithRecovery(),
	)

	// Load the city code map
	err := city.CityClient.LoadCodeMap()
	if err != nil {
		fmt.Printf("Error loading city code map: %v\n", err)
		return
	}

	// Add the forecast tool and its handler function
	s.AddTool(mcp.GetForecast(), mcp.ForecastCall)

	fmt.Println("Weather Demo started, waiting for requests...")
	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
