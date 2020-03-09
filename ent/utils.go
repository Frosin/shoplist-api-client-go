package ent

import (
	"context"
	"fmt"
)

func (c *Client) WithTx(ctx context.Context, fn func(tx *Tx) error) error {
	return withTx(ctx, c, fn)
}

func withTx(ctx context.Context, client *Client, fn func(tx *Tx) error) (err error) {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if v := recover(); v != nil {
			err = tx.Rollback()
		}
	}()

	if err := fn(tx); err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			err = fmt.Errorf("rolling back transaction: %w", rErr)
		}

		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}

func Rollback(tx *Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		return rerr
	}
	return err
}
