package main

import (
	"context"
	"fmt"
	"log"

	//"github.com/googlemaps/google-maps-services-go/maps"
	//"googlemaps.github.io/maps"
	"googlemaps.github.io/maps"
	//"github.com/googlemaps/google-maps-services-go/maps"
)

func main() {
	// Replace "YOUR_API_KEY" with your actual API key
	apiKey := "AIzaSyBuImtLS1A2933QJOZ-ScwDVeEfDYF9aSw"

	// Create a new client with your API key.
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating Google Maps client: %v", err)
	}

	// Geocoding request example: get coordinates for an address.
	address := "1600 Amphitheatre Parkway, Mountain View, CA"
	geocodeRequest := &maps.GeocodingRequest{
		Address: address,
	}

	// Perform the geocoding request.
	geocodeResult, err := client.Geocode(context.Background(), geocodeRequest)
	if err != nil {
		log.Fatalf("Error during geocoding request: %v", err)
	}

	// Print the results.
	for _, result := range geocodeResult {
		fmt.Printf("Formatted Address: %s\n", result.FormattedAddress)
		fmt.Printf("Latitude: %f\n", result.Geometry.Location.Lat)
		fmt.Printf("Longitude: %f\n", result.Geometry.Location.Lng)
		fmt.Println("-----")
	}
}
