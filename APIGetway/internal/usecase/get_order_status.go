package usecase

import (
	"context"

	orderdto "github.com/DencCPU/gRPCServices/APIGetway/internal/adapters/dto/order_service"
	"go.uber.org/zap"
)

func (s *Service) GetOrderStatus(ctx context.Context, input orderdto.GetInput) (orderdto.GetOutput, error) {
	ctx, span := s.tracer.Start(ctx, "Order status")
	defer span.End()

	output, err := s.order_client.GetStatus(ctx, input)
	if err != nil {
		s.logger.Error("error getting order status:",
			zap.String("spanID", span.SpanContext().SpanID().String()),
			zap.Error(err),
		)
		return orderdto.GetOutput{}, err
	}

	s.logger.Info("order status received",
		zap.String("spanID", span.SpanContext().SpanID().String()),
	)

	return output, nil
}
