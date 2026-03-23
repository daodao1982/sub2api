package repository

import "testing"

func TestShouldAutoDeleteFailedAccount(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want bool
	}{
		{name: "empty", msg: "", want: false},
		{name: "plain fail word", msg: "oauth fail while refreshing", want: true},
		{name: "capitalized fail", msg: "Fail", want: true},
		{name: "failed suffix", msg: "refresh token failed", want: true},
		{name: "failure noun", msg: "permanent failure from upstream", want: true},
		{name: "status equals fail", msg: "status=Fail", want: true},
		{name: "quoted fail", msg: "account state 'Fail' detected", want: true},
		{name: "chinese fail", msg: "账号失败，请删除", want: true},
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
