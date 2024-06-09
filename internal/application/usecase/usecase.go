package usecase

type UseCase[I any, O any] interface {
	Execute(I) (O, error)
}
