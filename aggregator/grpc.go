package main

import (
	"github.com/hamedblue1381/tolling/types"
)

type GRPCAggregatorServer struct {
	types.UnimplementedAggregatorServer
	svc Aggregator
}

func NewGRPCAggregatorServer(svc Aggregator) *GRPCAggregatorServer {
	return &GRPCAggregatorServer{
		svc: svc,
	}
}

func (s *GRPCAggregatorServer) AggregateDistance(req *types.AggregateRequest) error {
	distance := types.Distance{
		OBUID:     int(req.ObuID),
		Value:     req.Value,
		Timestamp: req.Timestamp,
	}
	return s.svc.AggregateDistance(distance)
}
