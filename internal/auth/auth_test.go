package auth

import (
	"strings"
	"testing"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name      string
		password  string
		expectErr bool
	}{
		{
			name:      "Valid password",
			password:  "Str0ng!Passw0rd",
			expectErr: false,
		},
		{
			name:      "Too short password",
			password:  "short",
			expectErr: true,
		},
		{
			name:      "Too long password",
			password:  strings.Repeat("a", 61),
			expectErr: true,
		},
		{
			name:      "Exactly 12 characters",
			password:  "1234567890ab",
			expectErr: false,
		},
		{
			name:      "Exactly 60 characters",
			password:  strings.Repeat("x", 60),
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashed, err := HashPassword(tt.password)

			if tt.expectErr {
				if err == nil {
					t.Errorf("Expected error for password %q, but got none", tt.password)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for password %q: %v", tt.password, err)
				} else {
					if hashed == "" {
						t.Error("Hashed password is empty")
					}
					if hashed == tt.password {
						t.Error("Hashed password should not be equal to the raw password")
					}
				}
			}
		})
	}
}
