package operator

import "testing"

func Test_Adder(t *testing.T) {
	type fields struct {
		x int
		y int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "10+6 gives 16",
			fields: fields{x: 10, y: 6},
			want:   16,
		},
		{
			name:   "20+40 gives 60",
			fields: fields{x: 20, y: 40},
			want:   60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &adder{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			if got := s.Add(); got != tt.want {
				t.Errorf("simpleClr.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}