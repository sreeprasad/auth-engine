package main

import (
	"auth-engine/auth"
	"fmt"
)

func main() {
	policies := []auth.Policy{
		{
			Statements: []auth.Statement{
				{
					Effect:    auth.Allow,
					Principal: []string{"user/eve"},
					Action:    []string{"read", "write"},
					Resource:  []string{"storage-service:::eve/*"},
					Condition: map[string]interface{}{
						"StringLike": map[string]string{
							"filename": "photo-*.jpg",
						},
					},
				},
			},
		},
		{
			Statements: []auth.Statement{
				{
					Effect:    auth.Deny,
					Principal: []string{"*"},
					Action:    []string{"delete"},
					Resource:  []string{"storage-service:::*/important-*"},
				},
			},
		},
	}

	// Let's tests some auth requests
	requests := []auth.Request{
		{
			Principal: "user/eve",
			Action:    "read",
			Resource:  "storage-service:::eve/photo-vacation.jpg",
			Context: map[string]interface{}{
				"filename": "photo-vacation.jpg",
			},
		},
		{
			Principal: "user/eve",
			Action:    "delete",
			Resource:  "storage-service:::eve/important-document.pdf",
		},
		{
			Principal: "user/alice",
			Action:    "read",
			Resource:  "storage-service:::eve/photo-vacation.jpg",
		},
	}

	for i, r := range requests {
		result := auth.Evaluate(policies, &r)
		fmt.Printf("Request %d: %s\n", i+1, result)
	}
}
