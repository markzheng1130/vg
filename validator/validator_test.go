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
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IsValidUUID(test.args.u); got != test.want {
				t.Errorf("IsValidUUID() = %v, want %v", got, test.want)
			}
		})
	}
}
