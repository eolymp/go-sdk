package ern

import (
	"context"
	"errors"
)

type Resolver interface {
	ResolveERN(context.Context, Resolver, ERN) (any, error)
}

type ResolverFunc func(context.Context, Resolver, ERN) (any, error)

func (f ResolverFunc) ResolveERN(ctx context.Context, r Resolver, ern ERN) (any, error) {
	return f(ctx, r, ern)
}

func NewResolver(rr ...Resolver) Resolver {
	return ResolverFunc(func(ctx context.Context, r Resolver, ern ERN) (any, error) {
		if !ern.Valid() {
			return nil, MalformedErr{}
		}

		for _, rs := range rr {
			resource, err := rs.ResolveERN(ctx, r, ern)
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
