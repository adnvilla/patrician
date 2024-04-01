package use_case

import "context"

type UseCase interface {
	Handle(ctx context.Context, in Input) (Output, error)
}

type Input interface{}
type Output interface{}
