package sqs_service

type Handler struct {
	useCase useCase
}

type useCase interface {
	// CreateCustomEvent(ctx context.Context, data entity.CustomEvent) error
}

func NewHandler(uc useCase) *Handler {
	return &Handler{
		useCase: uc,
	}
}

// // o que fazer com uma mensagem?
// func (h *Handler) HandleMessage(ctx context.Context, msg *sqs.Message) error {
// 	var data entity.CustomEvent

// 	err := json.Unmarshal([]byte(*msg.Body), &data)

// 	if err != nil {
// 		return fmt.Errorf("err: %w", err)
// 	}

// 	err = h.useCase.CreateCustomEvent(ctx, data)

// 	return err
// }
