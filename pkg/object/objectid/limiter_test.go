package objectid

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_ObjectID_Limiter(t *testing.T) {
	testCases := []struct {
		lis [][]string
		lim [][]string
	}{
		// Case 000
		{
			lis: [][]string{},
			lim: nil,
		},
		// Case 001
		{
			lis: [][]string{
				{"1", "2", "3"},
			},
			lim: [][]string{
				{"1", "2", "3"},
			},
		},
		// Case 002
		{
			lis: [][]string{
				{"1", "2", "3", "4", "5"},
			},
			lim: [][]string{
				{"1", "2", "3", "4", "5"},
			},
		},
		// Case 003
		{
			lis: [][]string{
				{"1", "2", "3", "4", "5", "6", "7"},
			},
			lim: [][]string{
				{"1", "2", "3", "4", "5"},
			},
		},
		// Case 004
		{
			lis: [][]string{
				{"1", "2", "3"},
				{"1", "2", "3"},
			},
			lim: [][]string{
				{"1", "2", "3"},
				{"1", "2"},
			},
		},
		// Case 005
		{
			lis: [][]string{
				{"1", "2"},
				{"1", "2", "3", "4", "5", "6"},
			},
			lim: [][]string{
				{"1", "2"},
				{"1", "2", "3"},
			},
		},
		// Case 006
		{
			lis: [][]string{
				{"1", "2", "3", "4", "5", "6"},
				{"1", "2"},
			},
			lim: [][]string{
				{"1", "2", "3", "4", "5"},
				{},
			},
		},
		// Case 007
		{
			lis: [][]string{
				{"1"},
				{"1", "2", "3"},
				{"1", "2"},
			},
			lim: [][]string{
				{"1"},
				{"1", "2", "3"},
				{"1"},
			},
		},
		// Case 008
		{
			lis: [][]string{
				{"1"},
				{"1", "2", "3"},
				{"1", "2"},
				{"1", "2", "3", "4", "5"},
				{"1"},
				{"1", "2"},
			},
			lim: [][]string{
				{"1"},
				{"1", "2", "3"},
				{"1"},
				{},
				{},
				{},
			},
		},
		// Case 009
		{
			lis: [][]string{
				{"1", "2"},
				{"1", "2", "3", "4", "5"},
				{"1"},
				{"1", "2", "3"},
				{"1"},
				{"1", "2"},
			},
			lim: [][]string{
				{"1", "2"},
				{"1", "2", "3"},
				{},
				{},
				{},
				{},
			},
		},
		// Case 010
		{
			lis: [][]string{
				{"1", "2", "3", "4", "5"},
				{"1", "2"},
				{"1", "2", "3"},
				{"1"},
				{"1"},
				{"1", "2"},
			},
			lim: [][]string{
				{"1", "2", "3", "4", "5"},
				{},
				{},
				{},
				{},
				{},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var l *Limiter
			{
				l = NewLimiter(5)
			}

			var lim [][]string
			for _, x := range tc.lis {
				lim = append(lim, x[:l.Limit(len(x))])
			}

			if !reflect.DeepEqual(lim, tc.lim) {
				t.Fatalf("expected %#v got %#v", tc.lim, lim)
			}
		})
	}
}
