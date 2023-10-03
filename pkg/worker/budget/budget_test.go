package budget

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Worker_Budget(t *testing.T) {
	testCases := []struct {
		lis [][]string
		cla [][]string
		bre bool
	}{
		// Case 000
		{
			lis: [][]string{},
			cla: nil,
			bre: false,
		},
		// Case 001
		{
			lis: [][]string{
				{"1", "2", "3"},
			},
			cla: [][]string{
				{"1", "2", "3"},
			},
			bre: false,
		},
		// Case 002
		{
			lis: [][]string{
				{"1", "2", "3", "4", "5"},
			},
			cla: [][]string{
				{"1", "2", "3", "4", "5"},
			},
			bre: false,
		},
		// Case 003
		{
			lis: [][]string{
				{"1", "2", "3", "4", "5", "6", "7"},
			},
			cla: [][]string{
				{"1", "2", "3", "4", "5"},
			},
			bre: true,
		},
		// Case 004
		{
			lis: [][]string{
				{"1", "2"},
				{"1", "2", "3"},
			},
			cla: [][]string{
				{"1", "2"},
				{"1", "2", "3"},
			},
			bre: false,
		},
		// Case 005
		{
			lis: [][]string{
				{"1", "2", "3"},
				{"1", "2", "3"},
			},
			cla: [][]string{
				{"1", "2", "3"},
				{"1", "2"},
			},
			bre: true,
		},
		// Case 006
		{
			lis: [][]string{
				{"1", "2"},
				{"1", "2", "3", "4", "5", "6"},
			},
			cla: [][]string{
				{"1", "2"},
				{"1", "2", "3"},
			},
			bre: true,
		},
		// Case 007
		{
			lis: [][]string{
				{"1", "2", "3", "4", "5", "6"},
				{"1", "2"},
			},
			cla: [][]string{
				{"1", "2", "3", "4", "5"},
				{},
			},
			bre: true,
		},
		// Case 008
		{
			lis: [][]string{
				{"1"},
				{"1", "2", "3"},
				{"1"},
			},
			cla: [][]string{
				{"1"},
				{"1", "2", "3"},
				{"1"},
			},
			bre: false,
		},
		// Case 009
		{
			lis: [][]string{
				{"1"},
				{"1", "2", "3"},
				{"1", "2"},
			},
			cla: [][]string{
				{"1"},
				{"1", "2", "3"},
				{"1"},
			},
			bre: true,
		},
		// Case 010
		{
			lis: [][]string{
				{"1"},
				{"1", "2", "3"},
				{"1", "2"},
				{"1", "2", "3", "4", "5"},
				{"1"},
				{"1", "2"},
			},
			cla: [][]string{
				{"1"},
				{"1", "2", "3"},
				{"1"},
				{},
				{},
				{},
			},
			bre: true,
		},
		// Case 011
		{
			lis: [][]string{
				{"1", "2"},
				{"1", "2", "3", "4", "5"},
				{"1"},
				{"1", "2", "3"},
				{"1"},
				{"1", "2"},
			},
			cla: [][]string{
				{"1", "2"},
				{"1", "2", "3"},
				{},
				{},
				{},
				{},
			},
			bre: true,
		},
		// Case 012
		{
			lis: [][]string{
				{"1", "2"},
				{"1", "2", "3"},
				{"1", "2", "3", "4", "5"},
				{"1"},
				{"1"},
				{"1", "2"},
			},
			cla: [][]string{
				{"1", "2"},
				{"1", "2", "3"},
				{},
				{},
				{},
				{},
			},
			bre: true,
		},
		// Case 013
		{
			lis: [][]string{
				{"1", "2", "3", "4", "5"},
				{"1", "2"},
				{"1", "2", "3"},
				{"1"},
				{"1"},
				{"1", "2"},
			},
			cla: [][]string{
				{"1", "2", "3", "4", "5"},
				{},
				{},
				{},
				{},
				{},
			},
			bre: true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var l *Budget
			{
				l = New(5)
			}

			var cla [][]string
			for _, x := range tc.lis {
				cla = append(cla, x[:l.Claim(len(x))])
			}

			if !reflect.DeepEqual(cla, tc.cla) {
				t.Fatalf("expected %#v got %#v", tc.cla, cla)
			}

			var bre bool
			{
				bre = l.Break()
			}

			if bre != tc.bre {
				t.Fatalf("expected %#v got %#v", tc.bre, bre)
			}
		})
	}
}
