package usecase

import (
	"context"

	orderdto "github.com/DencCPU/gRPCServices/APIGetway/internal/adapters/dto/order_service"
	orderdomain "github.com/DencCPU/gRPCServices/APIGetway/internal/domain/order"
	"go.uber.org/zap"
)

func (s *Service) CreateOrder(ctx context.Context, order orderdomain.OrderInfo) (orderdto.Output, error) {
	ctx, span := s.tracer.Start(ctx, "Create new order")
	defer span.End()

	output, err := s.order_client.CreateNewOrder(ctx, order)
	if err != nil {
		s.logger.Error("error creating a new order:",
			zap.String("spanID:", span.SpanContext().SpanID().String()),
			zap.Error(err),
		)
		return orderdto.Output{}, err
	}

	s.logger.Info("Order created",
		zap.String("spanID:", span.SpanContext().SpanID().String()),
	)

	return output, nil
}
