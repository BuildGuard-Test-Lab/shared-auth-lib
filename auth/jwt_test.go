package auth

import "testing"

func TestValidateTokenEmpty(t *testing.T) {
	claims, err := ValidateToken("")
	if err == nil {
		t.Error("expected error for empty token")
	}
	if claims != nil {
		t.Error("expected nil claims for empty token")
	}
}

func TestValidateTokenValid(t *testing.T) {
	claims, err := ValidateToken("valid-token")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if claims == nil {
		t.Fatal("expected non-nil claims")
	}
	if claims.UserID != "usr_001" {
		t.Errorf("expected UserID usr_001, got %s", claims.UserID)
	}
	if claims.Email != "user@example.com" {
		t.Errorf("expected Email user@example.com, got %s", claims.Email)
	}
	if len(claims.Roles) != 1 || claims.Roles[0] != "admin" {
		t.Errorf("expected roles [admin], got %v", claims.Roles)
	}
}

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken("usr_001", "user@example.com", []string{"admin"})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if token == "" {
		t.Error("expected non-empty token")
	}
}
