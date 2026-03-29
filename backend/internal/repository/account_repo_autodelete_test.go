package repository

import "testing"

func TestShouldAutoDeleteFailedAccount(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want bool
	}{
		{name: "empty", msg: "", want: false},
		{name: "generic fail should not delete", msg: "oauth fail while refreshing", want: false},
		{name: "generic failed should not delete", msg: "refresh token failed", want: false},
		{name: "generic failure should not delete", msg: "permanent failure from upstream", want: false},
		{name: "api returned 401", msg: "API returned 401: token invalid", want: true},
		{name: "json status 401", msg: `{"status": 401, "error": {"message":"bad token"}}`, want: true},
		{name: "token invalidated", msg: "scheduled health probe detected invalid account: token_invalidated", want: true},
		{name: "token revoked", msg: "authentication token has been invalidated; token_revoked", want: true},
		{name: "403 should not delete", msg: `Access forbidden (403): {"code":"invalid_workspace_selected"}`, want: false},
		{name: "429 should not delete", msg: `Rate limit exceeded (429): quota exhausted`, want: false},
		{name: "unrelated message", msg: "temporary overload cooldown", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shouldAutoDeleteFailedAccount(tt.msg)
			if got != tt.want {
				t.Fatalf("shouldAutoDeleteFailedAccount(%q) = %v, want %v", tt.msg, got, tt.want)
			}
		})
	}
}
