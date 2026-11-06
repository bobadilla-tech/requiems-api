package main

import (
	"context"
	"fmt"
	"os"
	"openapi"
)

func main() {
	cfg := openapi.NewConfiguration()
	client := openapi.NewAPIClient(cfg)

	_, r, err := client.AdviceAPI.V1TextAdviceGet(context.Background()).Execute()
	
	if err != nil {
		if r != nil && r.StatusCode == 401 {
			fmt.Println("✅ Go Client E2E OK (Expected 401)")
			os.Exit(0)
		}
		fmt.Printf("❌ Failed: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Println("❌ Failed: Expected 401 but got success")
	os.Exit(1)
}
