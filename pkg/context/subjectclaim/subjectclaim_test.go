package subjectclaim

import (
	"context"
	"testing"
)

func Test_Context_UserID(t *testing.T) {
	{
		s := FromContext(context.Background())
		if s != "" {
			t.Fatal("user must be empty")
		}
	}

	{
		s := FromContext(NewContext(context.Background(), "test"))
		if s != "test" {
			t.Fatal("user must be test")
		}
	}
}
