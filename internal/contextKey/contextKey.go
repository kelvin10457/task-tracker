package contextkey

import (
	"context"

	"github.com/kelvin10457/task-tracker/internal/db"
	"github.com/spf13/cobra"
)

type ctxKeyStore struct{}

func GetStore(cmd *cobra.Command) *db.Store {
	//this is a validation for knowing if exists the value for that key and if it's a db.Store
	val := cmd.Context().Value(ctxKeyStore{})
	if s, ok := val.(*db.Store); ok {
		return s
	}
	return nil
}

func SetStore(cmd *cobra.Command, store *db.Store) {
	ctx := context.WithValue(cmd.Context(), ctxKeyStore{}, store)
	cmd.SetContext(ctx)
}
