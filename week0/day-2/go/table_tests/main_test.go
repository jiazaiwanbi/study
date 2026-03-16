package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		x, y int
		want int
	}{
		{"1+2", 1, 2, 3},
		{"2+3", 2, 3, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.x, tt.y); got != tt.want {
				t.Fatalf("got=%d want=%d", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		x, y    int
		want    float64
		wantErr bool
	}{
		{"1/2", 1, 2, 0.5, false},
		{"2/0", 2, 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.x, tt.y)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err=%v wantErr=%v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Fatalf("got=%v want=%v", got, tt.want)
			}
		})
	}
}
