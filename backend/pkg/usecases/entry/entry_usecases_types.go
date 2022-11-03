package entry

import (
	"context"
	"time"
)

type Entry struct {
	ID           string
	Email        string
	Username     string
	Password     string
	CreationDate time.Time
	LastUpdate   time.Time
}

type CreateEntryUseCaseInput struct {
	ID        string `valid:"required,length(1|64)"`
	Email     string `valid:"required,length(1|64)"`
	Entryname string `valid:"required,length(1|64)"`
	Password  string `valid:"required,length(1|64)"`
}

type CreateEntryUseCaseOutput struct {
	Entry
}

type DescribeEntryUseCaseInput struct {
	ID string `valid:"required,length(1|64)"`
}

type DescribeEntryUseCaseOutput struct {
	Entry
}

type ListEntryUseCaseInput struct{}

type ListEntryUseCaseOutput struct {
	Items []Entry
}

type UpdateEntryUseCaseInput struct {
	ID        string `valid:"required,length(1|64)"`
	Email     string `valid:"required,length(1|64)"`
	Entryname string `valid:"required,length(1|64)"`
	Password  string `valid:"required,length(1|64)"`
}

type UpdateEntryUseCaseOutput struct {
	Entry
}

type RemoveEntryUseCaseInput struct {
	ID string `valid:"required,length(1|64)"`
}

type RemoveEntryUseCaseOutput struct {
	Entry
}

/* TYPES FOR STORAGE */

type AddInput struct {
	ID           string
	Email        string
	Entryname    string
	Password     string
	CreationDate time.Time
	LastUpdate   time.Time
}

type AddOutput struct {
	Entry
}

type GetInput struct {
	ID string
}

type GetOutput struct {
	Entry
}

type AllInput struct{}

type AllOutput struct {
	Items []Entry
}

type EditInput struct {
	ID         string
	Email      string
	Entryname  string
	Password   string
	LastUpdate time.Time
}

type EditOutput struct {
	Entry
}

type RemoveInput struct {
	ID string
}

type RemoveOutput struct {
	Entry
}

/* COLLECTION TYPE */

type EntryCollection []Entry

/* INTERFACES */

type EntryStorage interface {
	Add(ctx context.Context, input AddInput) (*AddOutput, error)
	Get(ctx context.Context, input GetInput) (*GetOutput, error)
	All(ctx context.Context, input AllInput) (*AllOutput, error)
	Edit(ctx context.Context, input EditInput) (*EditOutput, error)
	Remove(ctx context.Context, input RemoveInput) (*RemoveOutput, error)
}
