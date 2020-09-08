package dynamodb

import (
	"context"
	"testing"

	"github.com/matryer/is"
)

func TestXExistTable(t *testing.T) {
	is := is.NewRelaxed(t)
	ctx := context.Background()
	svc := getTestClient(t)

	t.Run("XExistTable", func(t *testing.T) {
		_ = svc.XDeleteTableFromName(ctx, testTableName)

		ok, err := svc.XExistTable(ctx, testTableName)
		is.True(!ok)
		is.NoErr(err)

		err = createTestTable(testTableName)
		is.NoErr(err)

		ok, err = svc.XExistTable(ctx, testTableName)
		is.True(ok)
		is.NoErr(err)
	})
}
