package customerdelete

import (
	"context"
)

type CustomerDelete interface {
	Handle(ctx context.Context, ID string) error
}
