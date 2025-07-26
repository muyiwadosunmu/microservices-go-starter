package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TripModel struct {
	ID       primitive.ObjectID `json:"id"`
	UserID   string             `json:"passenger"`
	Status   string             `json:"status"`
	RideFare *RideFareModel     `json:"ride_fare"`
	Duration float64            `json:"duration"`
}

type TripRepository interface {
	CreateTrip(context.Context, *TripModel) (*TripModel, error)
}

type TripService interface {
	CreateTrip(context.Context, RideFareModel) (*TripModel, error)
}
