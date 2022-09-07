package search

import (
	"github.com/zhouyusd/algorithm"
	"testing"
)

type linearArgs[T algorithm.Comparable] struct {
	a []T
	x T
}

type linearTestcase[T algorithm.Comparable] struct {
	name    string
	args    linearArgs[T]
	want    int
	wantErr bool
}

func LinearRunTests[T algorithm.Comparable](t *testing.T, cases []linearTestcase[T]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Linear(tt.args.a, tt.args.x)
			if (err != nil) != tt.wantErr {
				t.Errorf("Linear() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Linear() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinear(t *testing.T) {
	intCases := []linearTestcase[int]{
		{
			name: "Int找得到",
			args: linearArgs[int]{
				a: []int{1, 2, 3, 4, 5},
				x: 1,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "Int找不到",
			args: linearArgs[int]{
				a: []int{2, 3, 4, 5, 6},
				x: 1,
			},
			want:    -1,
			wantErr: true,
		},
	}
	LinearRunTests(t, intCases)
	stringCases := []linearTestcase[string]{
		{
			name: "String找得到",
			args: linearArgs[string]{
				a: []string{"1", "2", "3", "4", "5"},
				x: "1",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "String找不到",
			args: linearArgs[string]{
				a: []string{"2", "3", "4", "5", "6"},
				x: "1",
			},
			want:    -1,
			wantErr: true,
		},
	}
	LinearRunTests(t, stringCases)
}

type Linear1Args[T any] struct {
	a []T
	x T
}

type Linear1Testcase[T any] struct {
	name    string
	args    Linear1Args[T]
	want    int
	wantErr bool
}

func Linear1RunTests[T any](t *testing.T, cases []Linear1Testcase[T], cmp func(T, T) int8) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Linear1(tt.args.a, tt.args.x, cmp)
			if (err != nil) != tt.wantErr {
				t.Errorf("Linear1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Linear1() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinear1(t *testing.T) {
	type point struct {
		x, y int
	}
	pointCases := []Linear1Testcase[point]{
		{
			name: "point找得到",
			args: Linear1Args[point]{
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
			args: Linear1Args[point]{
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
	Linear1RunTests(t, pointCases, func(a, b point) int8 {
		if a.x == b.x {
			if a.y == b.y {
				return 0
			}
			return int8(a.y - b.y)
		}
		return int8(a.x - b.x)
	})
}
