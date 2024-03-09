package models

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

type User struct {
	ID            int    `json:"id" gorm:"primary_key"`
	Name          string `json:"name" `
	Email         string `json:"email" binding:"required,email"`
	ContactNumber string `json:"contactNumber" binding:"e164"`
	Role          string `json:"role"`
	Password      string `json:"password" binding:"required"`
	LibID         int    `json:"libID" `
}

type RequestEvents struct {
	ReqID        int `json:"id" gorm:"primary_key"`
	BookID       int `json:"bookid"`
	ReaderID     int	`json:"readerid"`
	RequestDate  time.Time	
	ApprovalDate time.Time	
	ApprovalID   int	
	RequestType  string
}

type Library struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

type IssueRegistry struct {
	IssueID            int `json:"id" gorm:"primary_key"`
	ISBN               int
	ReaderID           int
	IssueApprovedID    int
	IssueStatus        string
	IssueDate          time.Time
	ExpectedReturnDate time.Time
	ReturnDate         time.Time
	ReturnApprovedID   int
}

type BookInventory struct {
	ISBN            int    `json:"id" gorm:"primary_key"`
	LibID           int    `json:"libID,string"`
	Title           string `json:"title"`
	Authors         string `json:"authors"`
	Publisher       string `json:"publisher"`
	Version         int    `json:"version"`
	TotalCopies     int    `json:"totalCopies,string"`
	AvailableCopies int    `json:"availableCopies,string"`
}


