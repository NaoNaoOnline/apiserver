package subscriptionscrapehandler

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/go-cmp/cmp"
)

func Test_Worker_Handler_Subscription_Scrape_addStr(t *testing.T) {
	testCases := []struct {
		add [3]common.Address
		str []string
	}{
		// Case 000
		{
			add: [3]common.Address{},
			str: nil,
		},
		// Case 001
		{
			add: [3]common.Address{
				common.HexToAddress("foo"),
				common.HexToAddress("0x0"),
			},
			str: nil,
		},
		// Case 002
		{
			add: [3]common.Address{
				common.HexToAddress("foo"),
				common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
				common.HexToAddress("0x0"),
			},
			str: []string{
				"0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
			},
		},
		// Case 003
		{
			add: [3]common.Address{
				common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
				common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
				common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
			},
			str: []string{
				"0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
				"0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
				"0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			str := addStr(tc.add)
			if !reflect.DeepEqual(str, tc.str) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.str, str))
			}
		})
	}
}
