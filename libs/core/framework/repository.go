package framework

import (
	"context"
)

type RawRepository[T Entity] interface {
	Create(ctx context.Context, entity *T) error
}

type RepositoryHooks[T Entity] struct {
	PreCreate  []func(ctx context.Context, entity *T) error
	PostCreate []func(ctx context.Context, entity *T) error
}

type Repository[T Entity] struct {
	repo  RawRepository[T]
	hooks RepositoryHooks[T]
}

func NewRepository[T Entity](repo RawRepository[T]) *Repository[T] {
	return &Repository[T]{
		repo:  repo,
		hooks: RepositoryHooks[T]{},
	}
}

func (r *Repository[T]) AddPreCreateHook(hook func(ctx context.Context, entity *T) error) *Repository[T] {
	r.hooks.PreCreate = append(r.hooks.PreCreate, hook)
	return r
}

func (r *Repository[T]) AddPostCreateHook(hook func(ctx context.Context, entity *T) error) *Repository[T] {
	r.hooks.PostCreate = append(r.hooks.PostCreate, hook)
	return r
}

func (r *Repository[T]) Create(ctx context.Context, entity *T) error {
	// Esegui pre-hooks specifici per create
	for _, hook := range r.hooks.PreCreate {
		if err := hook(ctx, entity); err != nil {
			return err
		}
	}

	err := r.repo.Create(ctx, entity)

	if err == nil {
		// Esegui post-hooks specifici per create
		for _, hook := range r.hooks.PostCreate {
			if err := hook(ctx, entity); err != nil {
				return err
			}
		}
	}

	return err
}
