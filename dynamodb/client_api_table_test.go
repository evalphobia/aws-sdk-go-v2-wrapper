package dynamodb

import (
	"context"
	"testing"

	"github.com/matryer/is"
)

func TestExistTable(t *testing.T) {
	is := is.NewRelaxed(t)
	ctx := context.Background()
	svc := getTestClient(t)

	t.Run("ExistTable", func(t *testing.T) {
		_ = svc.DeleteTableFromName(ctx, testTableName)

		ok, err := svc.ExistTable(ctx, testTableName)
		is.True(!ok)
		is.NoErr(err)

		err = createTestTable(testTableName)
		is.NoErr(err)

		ok, err = svc.ExistTable(ctx, testTableName)
		is.True(ok)
		is.NoErr(err)
	})
}
