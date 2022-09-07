package search

import (
	"github.com/zhouyusd/algorithm"
	"testing"
)

type binaryArgs[T algorithm.Comparable] struct {
	a []T
	x T
}

type binaryTestcase[T algorithm.Comparable] struct {
	name    string
	args    binaryArgs[T]
	want    int
	wantErr bool
}

func BinaryRunTests[T algorithm.Comparable](t *testing.T, cases []binaryTestcase[T]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Binary(tt.args.a, tt.args.x)
			if (err != nil) != tt.wantErr {
				t.Errorf("Binary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Binary() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinary(t *testing.T) {
	intCases := []binaryTestcase[int]{
		{
			name: "Int找得到",
			args: binaryArgs[int]{
				a: []int{1, 2, 3, 4, 5},
				x: 1,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "Int找不到",
			args: binaryArgs[int]{
				a: []int{2, 3, 4, 5, 6},
				x: 1,
			},
			want:    -1,
			wantErr: true,
		},
	}
	BinaryRunTests(t, intCases)
	stringCases := []binaryTestcase[string]{
		{
			name: "String找得到",
			args: binaryArgs[string]{
				a: []string{"1", "2", "3", "4", "5"},
				x: "1",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "String找不到",
			args: binaryArgs[string]{
				a: []string{"2", "3", "4", "5", "6"},
				x: "1",
			},
			want:    -1,
			wantErr: true,
		},
	}
	BinaryRunTests(t, stringCases)
}

type Binary1Args[T any] struct {
	a []T
	x T
}

type Binary1Testcase[T any] struct {
	name    string
	args    Binary1Args[T]
	want    int
	wantErr bool
}

func Binary1RunTests[T any](t *testing.T, cases []Binary1Testcase[T], cmp func(T, T) int8) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Binary1(tt.args.a, tt.args.x, cmp)
			if (err != nil) != tt.wantErr {
				t.Errorf("Binary1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Binary1() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinary1(t *testing.T) {
	type point struct {
		x, y int
	}
	pointCases := []Binary1Testcase[point]{
		{
			name: "point找得到",
			args: Binary1Args[point]{
				a: []point{
					{
						x: 0,
						y: 0,
					},
					{
						x: 1,
						y: 1,
					},
				},
				x: point{
					x: 0,
					y: 0,
				},
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "point找不到",
			args: Binary1Args[point]{
				a: []point{
					{
						x: 0,
						y: 1,
					},
					{
						x: 1,
						y: 1,
					},
				},
				x: point{
					x: 0,
					y: 0,
				},
			},
			want:    -1,
			wantErr: true,
		},
	}
	Binary1RunTests(t, pointCases, func(a, b point) int8 {
		if a.x == b.x {
			if a.y == b.y {
				return 0
			}
			return int8(a.y - b.y)
		}
		return int8(a.x - b.x)
	})
}
