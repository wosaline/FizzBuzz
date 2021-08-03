//three integers int1, int2 and limit, and two strings str1and str2

package service

import "testing"

func TestFizzBuzz(t *testing.T) {
	type args struct {
		limit     int
		multiple1 int
		multiple2 int
		str1      string
		str2      string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Simple test OK", args{5, 2, 3, "fizz", "buzz"}, "1,fizz,buzz,fizz,5", false},
		{"Should return Error when limit is inferior to 1 or negative", args{0, 2, 3, "fizz", "buzz"}, "", true},
		{"Should succeed with bigger scale sample", args{30, 3, 5, "fizz", "buzz"}, "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,17,fizz,19,buzz,fizz,22,23,fizz,buzz,26,fizz,28,29,fizzbuzz", false},
		{"Should return Error when one of the multiples is inferior to 1 or negative", args{4, 0, 0, "fizz", "buzz"}, "", true},
		{"Should succeed when str1 or str2 is blank", args{5, 2, 3, "", ""}, "1,,,,5", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FizzBuzz(tt.args.limit, tt.args.multiple1, tt.args.multiple2, tt.args.str1, tt.args.str2)
			if (err != nil) != tt.wantErr {
				t.Errorf("FizzBuzz() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
