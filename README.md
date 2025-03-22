# Auth Engine

A Go implementation of a policy-based authorization engine inspired by the paper [Formally Verified Cloud-Scale Authorization](https://www.amazon.science/publications/formally-verified-cloud-scale-authorization).

## Overview

This project provides a simple implementation of an authorization engine that evaluates whether a request should be allowed or denied based on policy statements. It demonstrates core concepts from cloud-scale authorization systems, including:

- Policy evaluation with "Deny Trumps Allow" semantics
- Wildcard pattern matching for flexible resource and principal targeting
- Separation of concerns between policy representation and evaluation logic

## Features

- Define policies with Allow/Deny effects
- Create statements with principal, action, resource, and condition elements
- Evaluate requests against policies with proper precedence rules
- Pattern matching with wildcards to support flexible policy definitions

## Usage

The engine allows define authorization policies and evaluate access requests against them:

```go
// How to define policies
policies := []auth.Policy{
    {
        Statements: []auth.Statement{
            {
                Effect:    auth.Allow,
                Principal: []string{"user/alice"},
                Action:    []string{"read", "list*"},
                Resource:  []string{"storage-service:::folder/*"},
            },
        },
    },
}

// How to evaluate a request
request := auth.Request{
    Principal: "user/alice",
    Action:    "read",
    Resource:  "storage-service:::folder/document.txt",
}

result := auth.Evaluate(policies, &request)
// Here the result will be AllowAccess, ExplicitDeny, or ImplicitDeny
```

## Getting Started
```bash
# Clone the repository
git clone https://github.com/sreeprasad/auth-engine.git

# Build the project
make build

# Run tests
make test
```

## Project Structure
```bash 
auth/ - Core authorization engine package
  types.go - Data model definitions
  evaluation.go - Policy evaluation logic
  matching.go - Wildcard pattern matching impl
main.go - Example usage
```


## License
MIT
