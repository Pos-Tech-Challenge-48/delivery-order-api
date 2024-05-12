package productdelete

import "context"

type ProductDelete interface {
	Delete(ctx context.Context, producID string) error
}
