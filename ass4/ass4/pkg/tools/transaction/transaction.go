package transaction

import (
	context2 "context"
	"github.com/jackc/pgx/v4"

	log "architecture_go/pkg/type/logger"
)

func Finish(ctx context2.Context, tx pgx.Tx, err error) error {
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return log.ErrorWithContext(ctx, rollbackErr)
		}
		return err
	} else {
		if commitErr := tx.Commit(ctx); commitErr != nil {
			return log.ErrorWithContext(ctx, commitErr)
		}
		return nil
	}
}
