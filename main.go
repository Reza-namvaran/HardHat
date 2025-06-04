package main

import (
	"fmt"
	"os"
	"net/http"
	"time"
)

var securityHeaders = []string{
    "Content-Security-Policy",
    "Strict-Transport-Security",
    "X-Frame-Options",
    "X-Content-Type-Options",
    "X-XSS-Protection",
    "Referrer-Policy",
    "Permissions-Policy",
}


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <url>");
		return;
	}

	url := os.Args[1];
	
	client := &http.Client{Timeout: 5 * time.Second};
	resp, err := client.Get(url);

	if err != nil {
		fmt.Printf("Failed to Fetech URL %v\n", err);
		return;
	}

	defer resp.Body.Close();


	fmt.Println("\nSecurity Header Check:")
	for _, header := range securityHeaders {
		if val := resp.Header.Get(header); val != "" {
			fmt.Printf("✅ %s: Present\n", header)
		} else {
			fmt.Printf("❌ %s: Missing\n", header)
		}
	}
}