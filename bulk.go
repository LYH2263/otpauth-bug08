package otpauth

import "context"

type Pair struct {
	ID   string
	Code string
}

func BulkVerify(ctx context.Context, v *Verifier, pairs []Pair) error {
	for _, p := range pairs {
		// Honor caller cancellation before each item so a cancelled
		// batch (short-timeout import, client disconnect) stops the loop
		// promptly instead of grinding through the whole batch.
		if err := ctx.Err(); err != nil {
			return err
		}
		if err := runVerify(ctx, v, p.ID, p.Code); err != nil {
			return err
		}
	}
	return nil
}
