package auth

// Effect represents whether an action is allowed or denied
type Effect string

const (
	Allow Effect = "Allow"
	Deny  Effect = "Deny"
)

// Answer represents the result of an authorization decision
type Answer string

const (
	ExplicitDeny Answer = "ExplicitDeny"
	AllowAccess  Answer = "Allow"
	ImplicitDeny Answer = "ImplicitDeny"
)

// Statement represents a single authorization rule
type Statement struct {
	Effect    Effect
	Principal []string
	Action    []string
	Resource  []string
	Condition map[string]interface{}
}

// Policy represents a collection of statements
type Policy struct {
	Statements []Statement
}

// Request represents an authorization request
type Request struct {
	Principal string
	Action    string
	Resource  string
	Context   map[string]interface{}
}
