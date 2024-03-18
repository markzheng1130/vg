package validator

import "testing"

func TestIsValidUUID(t *testing.T) {
	type args struct {
		u string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// test cases
		{
			"caseReturnTrue",
			args{
				u: "a0a8aea5-cc40-4293-b1f3-c3bc4e53d941",
			},
			true,
		},
		{
			"caseReturnFalse",
			args{
				u: "a0a8aea5-cc40-4293-b1f3-",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidUUID(tt.args.u); got != tt.want {
				t.Errorf("IsValidUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}
