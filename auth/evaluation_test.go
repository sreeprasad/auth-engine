package auth

import (
	"testing"
)

func TestStatementMatches(t *testing.T) {
	statement := Statement{
		Effect:    Allow,
		Principal: []string{"user/alice", "user/*admin*"},
		Action:    []string{"read", "list*"},
		Resource:  []string{"resource/folder/*"},
	}

	tests := []struct {
		name    string
		request Request
		want    bool
	}{
		{
			name: "Full match",
			request: Request{
				Principal: "user/alice",
				Action:    "read",
				Resource:  "resource/folder/document.txt",
			},
			want: true,
		},
		{
			name: "Principal pattern match",
			request: Request{
				Principal: "user/superadmin",
				Action:    "read",
				Resource:  "resource/folder/document.txt",
			},
			want: true,
		},
		{
			name: "Action pattern match",
			request: Request{
				Principal: "user/alice",
				Action:    "listFiles",
				Resource:  "resource/folder/document.txt",
			},
			want: true,
		},
		{
			name: "No principal match",
			request: Request{
				Principal: "user/bob",
				Action:    "read",
				Resource:  "resource/folder/document.txt",
			},
			want: false,
		},
		{
			name: "No action match",
			request: Request{
				Principal: "user/alice",
				Action:    "write",
				Resource:  "resource/folder/document.txt",
			},
			want: false,
		},
		{
			name: "No resource match",
			request: Request{
				Principal: "user/alice",
				Action:    "read",
				Resource:  "other/folder/document.txt",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := statement.Matches(&tt.request); got != tt.want {
				t.Errorf("Statement.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvaluate(t *testing.T) {
	policies := []Policy{
		{
			Statements: []Statement{
				{
					Effect:    Allow,
					Principal: []string{"user/eve"},
					Action:    []string{"read", "write"},
					Resource:  []string{"storage-service:::eve/*"},
				},
			},
		},
		{
			Statements: []Statement{
				{
					Effect:    Deny,
					Principal: []string{"*"},
					Action:    []string{"delete"},
					Resource:  []string{"storage-service:::*/important-*"},
				},
			},
		},
	}

	tests := []struct {
		name    string
		request Request
		want    Answer
	}{
		{
			name: "Allowed request",
			request: Request{
				Principal: "user/eve",
				Action:    "read",
				Resource:  "storage-service:::eve/document.txt",
			},
			want: AllowAccess,
		},
		{
			name: "Explicitly denied request",
			request: Request{
				Principal: "user/eve",
				Action:    "delete",
				Resource:  "storage-service:::eve/important-document.pdf",
			},
			want: ExplicitDeny,
		},
		{
			name: "Implicitly denied request - wrong principal",
			request: Request{
				Principal: "user/alice",
				Action:    "read",
				Resource:  "storage-service:::eve/document.txt",
			},
			want: ImplicitDeny,
		},
		{
			name: "Deny trumps allow",
			request: Request{
				Principal: "user/eve",
				Action:    "delete",
				Resource:  "storage-service:::eve/important-document.pdf",
			},
			want: ExplicitDeny,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Evaluate(policies, &tt.request); got != tt.want {
				t.Errorf("Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
