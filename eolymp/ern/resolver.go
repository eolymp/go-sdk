package ern

import (
	"context"
	"errors"
)

type Resolver interface {
	ResolveERN(ctx context.Context, ern ERN) (any, error)
}

type ResolverFunc func(ctx context.Context, ern ERN) (any, error)

func (f ResolverFunc) ResolveERN(ctx context.Context, ern ERN) (any, error) {
	return f(ctx, ern)
}

func NewResolver(rr ...Resolver) Resolver {
	return ResolverFunc(func(ctx context.Context, ern ERN) (any, error) {
		for _, r := range rr {
			resource, err := r.ResolveERN(ctx, ern)
			if errors.Is(err, InvalidErr{}) {
				continue
			}

			if err != nil {
				return nil, err
			}

			return resource, nil
		}

		return nil, InvalidErr{}
	})
}
