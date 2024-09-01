package model

type WorkerManagerInterface interface {
	Name() string
	Run() error
	Clear() error
	Shutdown() error
	Health() bool
}
