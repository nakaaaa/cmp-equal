package cmpequal

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
)

func TestSimpleEqual(t *testing.T) {
	sum := 2

	tests := []struct {
		name     string
		expected int
		wantErr  bool
	}{
		{
			name:     "ok",
			expected: 3,
			wantErr:  false,
		},
		{
			name:     "ng",
			expected: 10,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := testSum(sum)
			fmt.Println(actual)

			if !tt.wantErr {
				Equal(t, tt.expected, actual)
				return
			}
			assert.NotEqual(t, tt.expected, actual)
		})
	}
}

type TestStruct struct {
	ID        uint64
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func TestStructEqual(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name     string
		expected *TestStruct
		actual   *TestStruct
		opts     []cmp.Option
		wantErr  bool
	}{
		{
			name: "ok",
			expected: &TestStruct{
				ID:        1,
				Name:      "test",
				Password:  "password",
				CreatedAt: now,
				UpdatedAt: now,
			},
			actual: &TestStruct{
				ID:       1,
				Name:     "test",
				Password: "password",
			},
			opts: []cmp.Option{
				cmpopts.IgnoreFields(TestStruct{}, "CreatedAt", "UpdatedAt"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Equal(t, tt.expected, tt.actual, tt.opts...)
		})
	}
}

func testSum(sum int) int {
	return sum + 1
}
