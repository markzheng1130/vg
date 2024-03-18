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
				u: "11111111-1111-1111-1111-111111111111",
			},
			true,
		},
		{
			"caseReturnFalse",
			args{
				u: "abc",
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
