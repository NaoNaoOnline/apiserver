package keyfmt

import (
	"fmt"
	"testing"
)

func Test_KeyFmt_Index(t *testing.T) {
	testCases := []struct {
		str string
		trm string
	}{
		// Case 000
		{
			str: "",
			trm: "",
		},
		// Case 001
		{
			str: "foo",
			trm: "foo",
		},
		// Case 002
		{
			str: "hello world",
			trm: "hello%20world",
		},
		// Case 003
		{
			str: " hello    world  ",
			trm: "hello%20world",
		},
		// Case 004
		{
			str: " \t\n Hello,     Gophers \n\t\r\n",
			trm: "hello%2C%20gophers",
		},
		// Case 005
		{
			str: " 030\t\naka     030 \n\t\r\n",
			trm: "030%20aka%20030",
		},
		// Case 006
		{
			str: "MEV",
			trm: "mev",
		},
		// Case 007
		{
			str: " MEV",
			trm: "mev",
		},
		// Case 008
		{
			str: "DeFi",
			trm: "defi",
		},
		// Case 009
		{
			str: "DeFi ",
			trm: "defi",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			trm := Index(tc.str)
			if trm != tc.trm {
				t.Fatalf("expected %#v got %#v", tc.trm, trm)
			}
		})
	}
}

func Test_KeyFmt_Label(t *testing.T) {
	testCases := []struct {
		str string
		trm string
	}{
		// Case 000
		{
			str: "",
			trm: "",
		},
		// Case 001
		{
			str: "foo",
			trm: "foo",
		},
		// Case 002
		{
			str: "hello world",
			trm: "hello%20world",
		},
		// Case 003
		{
			str: " hello    world  ",
			trm: "hello%20world",
		},
		// Case 004
		{
			str: " \t\n Hello,     Gophers \n\t\r\n",
			trm: "Hello%2C%20Gophers",
		},
		// Case 005
		{
			str: " 030\t\naka     030 \n\t\r\n",
			trm: "030%20aka%20030",
		},
		// Case 006
		{
			str: "MEV",
			trm: "MEV",
		},
		// Case 007
		{
			str: " MEV",
			trm: "MEV",
		},
		// Case 008
		{
			str: "DeFi",
			trm: "DeFi",
		},
		// Case 009
		{
			str: "DeFi ",
			trm: "DeFi",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			trm := Label(tc.str)
			if trm != tc.trm {
				t.Fatalf("expected %#v got %#v", tc.trm, trm)
			}
		})
	}
}
