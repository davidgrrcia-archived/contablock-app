package entryInputHttpAdapter

import (
	"time"
)

type Entry struct {
	ID           string    `json:"id"`
	Entryname    string    `json:"entryname"`
	Password     string    `json:"password"`
	Email        string    `json:"email"`
	CreationDate time.Time `json:"creation-date"`
	LastUpdate   time.Time `json:"last-update"`
}

type CreateEntryRequest struct {
	ID        string `json:"id"`
	Entryname string `json:"entryname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}
type CreateEntryResponse struct {
	Entry
}
type CreateEntryResponseWrapper struct {
	CreateEntryResponse CreateEntryResponse
}

type DescribeEntryRequest struct {
	ID string `json:"id"`
}
type DescribeEntryResponse struct {
	Entry
}
type DescribeEntryResponseWrapper struct {
	DescribeEntryResponse DescribeEntryResponse
}

type ListEntryRequest struct{}

type ListEntryResponse struct {
	Items []DescribeEntryResponse `json:"items"`
}
type ListEntryResponseWrapper struct {
	ListEntryResponse ListEntryResponse
}

type UpdateEntryRequest struct {
	ID        string `json:"id"`
	Entryname string `json:"entryname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}
type UpdateEntryResponse struct {
	Entry
}
type UpdateEntryResponseWrapper struct {
	UpdateEntryResponse UpdateEntryResponse
}

type DeleteEntryRequest struct {
	ID string `json:"id"`
}
type DeleteEntryResponse struct {
	Entry
}
type DeleteEntryResponseWrapper struct {
	DeleteEntryResponse DeleteEntryResponse
}
