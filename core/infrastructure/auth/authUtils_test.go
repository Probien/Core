package auth

import "testing"

func TestAuthoritiesChecker(t *testing.T) {

	expected := true
	got := checkAuthorities([]string{"---", "", "ROLE_TEST", "ROLE_EMPLOYEE"}, &AuthCustomClaims{Roles: map[string]string{"role_0": "FAKER", "role_1": "FAKER", "role_2": "ROLE_EMPLOYEE", "role_3": "FAKER"}})

	if expected != got {
		t.Errorf("Expected %t, got %t", expected, got)
	} else {
		t.Logf("Expected %t, got %t", expected, got)
	}

}
