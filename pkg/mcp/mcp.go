package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	cityMap "mcp-weather/internal/city"
	"mcp-weather/internal/weather"

	"github.com/mark3labs/mcp-go/mcp"
)

type ForecastResult struct {
	Name        string `json:"Name"`
	Temperature string `json:"Temperature"`
	Wind        string `json:"Wind"`
	Forecast    string `json:"Forecast"`
	Date        string `json:"Date"`
}

// GetForecast Add forecast tool
func GetForecast() mcp.Tool {
	weatherTool := mcp.NewTool("get_weather_forecast",
		mcp.WithDescription("Get weather forecast for a city"),
		mcp.WithString("city",
			mcp.Required(),
			mcp.Description("city name"),
		),
		mcp.WithString("type",
			mcp.Required(),
			mcp.Description("get weather forecast info base/all"),
			mcp.Enum("base", "all"),
		),
	)
	return weatherTool
}

func ForecastCall(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// get city name
	city := request.Params.Arguments["city"].(string)
	if len(city) == 0 {
		return nil, fmt.Errorf("city name is empty")
	}

	// get info type
	infoType := request.Params.Arguments["type"].(string)
	if len(infoType) == 0 {
		return nil, fmt.Errorf("info type is empty")
	}

	adcode, exists := cityMap.CityClient.GetAdcode(city)
	if !exists {
		return nil, fmt.Errorf("city %s not found, code %s", city, adcode)
	}

	// get weather forecast
	weatherInfo, err := weather.FetchWeatherData(adcode, infoType)
	if err != nil {
		return nil, err
	}
	// parse weather info
	forecast := make([]ForecastResult, 0)

	if infoType == "base" {
		// parse base info
		var baseInfo weather.BaseInfo
		if err := json.Unmarshal([]byte(weatherInfo), &baseInfo); err != nil {
			return nil, fmt.Errorf("parse base info failed: %v", err)
		}
		for _, forecastInfo := range baseInfo.Lives {
			// append forecast info
			forecast = append(forecast, ForecastResult{
				Name:        forecastInfo.City,
				Temperature: forecastInfo.Temperature,
				Wind:        forecastInfo.Windpower,
				Forecast:    forecastInfo.Weather,
				Date:        forecastInfo.Reporttime,
			})
		}
	} else if infoType == "all" {
		// parse all info
		var allInfo weather.AllInfo
		if err := json.Unmarshal([]byte(weatherInfo), &allInfo); err != nil {
			return nil, fmt.Errorf("parse all info failed: %v", err)
		}
		for _, forecastInfo := range allInfo.Forecasts {
			for _, cast := range forecastInfo.Casts {
				forecast = append(forecast, ForecastResult{
					Name:        forecastInfo.City,
					Temperature: cast.Daytemp,
					Wind:        cast.Daywind,
					Forecast:    cast.Dayweather,
					Date:        cast.Date,
				})
			}
		}
	}

	forecastResponse, err := json.Marshal(&forecast)
	if err != nil {
		return nil, fmt.Errorf("error marshalling forecast: %w", err)
	}

	return mcp.NewToolResultText(string(forecastResponse)), nil
}
