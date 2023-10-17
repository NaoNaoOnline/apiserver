package policycache

import (
	"testing"

	fuzz "github.com/google/gofuzz"
)

func Test_Cache_Policy_Memory_Buffer_Fuzz(t *testing.T) {
	var pol Interface
	{
		pol = Fake()
	}

	var fuz *fuzz.Fuzzer
	{
		fuz = fuzz.New()
	}

	for i := 0; i < 1000; i++ {
		var inp []*Record
		{
			inp = []*Record{}
		}

		{
			fuz.Fuzz(&inp)
		}

		{
			_ = pol.Buffer(inp)
		}
	}
}
