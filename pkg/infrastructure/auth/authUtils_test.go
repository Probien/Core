package auth

import (
	"testing"
)

func TestAuthoritiesChecker(t *testing.T) {

	type errorCaseTest struct {
		description      string
		rolesInput       CustomClaims
		authoritiesInput []string
		expected         bool
	}

	tests := []errorCaseTest{
		{
			description: "invalid roles",
			rolesInput: CustomClaims{
				Roles: map[string]string{"role_0": "FAKER", "role_1": "FAKER", "role_2": "FAKER", "role_3": "FAKER"},
			},
			authoritiesInput: []string{"ROLE_ADMIN", "ROLE_EMPLOYEE"},
			expected:         false,
		},
		{
			description: "invalid empty roles",
			rolesInput: CustomClaims{
				Roles: map[string]string{},
			},
			authoritiesInput: []string{"ROLE_ADMIN"},
			expected:         false,
		},
		{
			description: "invalid role matches",
			rolesInput: CustomClaims{
				Roles: map[string]string{"role_0": "ROLE_SUPERVISOR", "role_1": "ROLE_GENERAL"},
			},
			authoritiesInput: []string{"ROLE_ADMIN", "ROLE_EMPLOYEE"},
			expected:         false,
		},
		{
			description: "valid roles",
			rolesInput: CustomClaims{
				Roles: map[string]string{"role_0": "ROLE_EMPLOYEE"},
			},
			authoritiesInput: []string{"ROLE_ADMIN", "ROLE_EMPLOYEE", "ROLE_GENERAL"},
			expected:         true,
		},
	}

	for _, v := range tests {
		t.Run(v.description, func(t *testing.T) {
			got := checkAuthorities(v.authoritiesInput, &v.rolesInput)
			if v.expected != got {
				t.Errorf("Expected %t, got %t", v.expected, got)
			} else {
				t.Logf("Expected %t, got %t", v.expected, got)
			}
		})
	}

}

func TestEncryptPassword(t *testing.T) {
	result := make(chan []byte, 1)
	EncryptPassword([]byte("kmzwa8awaa"), result)
	encryptedPassword := <-result

	if len(encryptedPassword) > 0 {
		t.Logf("encrypted password is: %v", string(encryptedPassword))
	} else {
		t.Fatalf("error encripting password")
	}
}
