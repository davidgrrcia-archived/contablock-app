package user

import (
	"context"
	"time"
)

type User struct {
	Project      string
	ID           string
	Name         string
	Email        string
	Password     string
	Role         string
	EthAddress   string
	CreationDate time.Time
	LastUpdate   time.Time
}

type DescribeUserInput struct {
	Project string `valid:"required,length(1|64)"`
	ID      string `valid:"required,length(1|64)"`
}

type DescribeUserOutput struct {
	User User
}

type ListUserInput struct {
	Project string `valid:"required,length(1|64)"`
}

type ListUserOutput struct {
	Items []User
}

type CreateUserInput struct {
	Project    string `valid:"required,length(1|64)"`
	ID         string `valid:"required,length(1|64)"`
	Name       string `valid:"required,length(1|256)"`
	Email      string `valid:"required,email,length(1|256)"`
	Password   string `valid:"required,length(1|256)"`
	Role       string `valid:"required,length(1|64)"`
	EthAddress string `valid:"required,length(1|42)"`
}

type CreateUserOutput struct {
	User User
}

type LoginUserInput struct {
	Project  string `valid:"required,length(1|64)"`
	Email    string `valid:"required,email,length(1|256)"`
	Password string `valid:"required,length(1|256)"`
}

type LoginUserOutput struct {
	User User
}

type Filters struct {
	Email    *string
	Password *string
}

type UserStorage interface {
	Get(ctx context.Context, project string, id string) (*UserStorageData, error)
	All(ctx context.Context, project string, filters Filters) (*UserStorageCollectionData, error)
	Edit(ctx context.Context, input UserStorageData) (*UserStorageData, error)
	Add(ctx context.Context, input UserStorageData) (*UserStorageData, error)
}

type UserStorageData struct {
	Project      string
	ID           string
	Name         string
	Email        string
	Password     string
	Role         string
	EthAddress   string
	CreationDate time.Time
	LastUpdate   time.Time
}

type UserStorageCollectionData struct {
	Items []UserStorageData
}
