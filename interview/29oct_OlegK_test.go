package main

import "testing"

func Test_getCapitalLetterWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"empty",
			args{s: "    "},
			"",
		},
		{
			"dots",
			args{s: "....."},
			"",
		},
		{
			"noCapitals",
			args{s: "asajkd ajdsa , asksa. asaf"},
			"",
		},
		{
			"allCapitals",
			args{s: "Aroirjari.. Aigjsig, ADKSKD,, "},
			"Aroirjari Aigjsig ADKSKD",
		},
		{
			"simple",
			args{s: "Aroirjari. aigjsig, ADKSKD"},
			"Aroirjari ADKSKD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCapitalLetterWords(tt.args.s); got != tt.want {
				t.Errorf("getCapitalLetterWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
