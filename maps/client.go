package maps

import "context"

type Client interface {
	Geocoder(ctx context.Context, keyword string) (*Geocode, error)
}
