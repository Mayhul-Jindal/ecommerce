// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package sqlc

import (
	"time"

	"github.com/google/uuid"
)

type Address struct {
	UserID      int64  `json:"user_id"`
	AddressLine string `json:"address_line"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
}

type Book struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	TagsArray   []int32   `json:"tags_array"`
	Price       int32     `json:"price"`
	Quantity    int32     `json:"quantity"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type Cart struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CartItem struct {
	ID        int64     `json:"id"`
	CartID    int64     `json:"cart_id"`
	BookID    int64     `json:"book_id"`
	Quantity  int32     `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}

type Order struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	TotalMoney  int32     `json:"total_money"`
	OrderStatus bool      `json:"order_status"`
	CreatedAt   time.Time `json:"created_at"`
}

type OrderLine struct {
	ID        int64     `json:"id"`
	BookID    int64     `json:"book_id"`
	OrderID   int64     `json:"order_id"`
	Quantity  int32     `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}

type Review struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	BookID    int64     `json:"book_id"`
	Rating    int32     `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	UserID       int64     `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Tag struct {
	ID      int32  `json:"id"`
	TagName string `json:"tag_name"`
}

type User struct {
	ID                int64     `json:"id"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	HashedPassword    string    `json:"hashed_password"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	IsAdmin           bool      `json:"is_admin"`
	IsActive          bool      `json:"is_active"`
	DeactivatedAt     time.Time `json:"deactivated_at"`
	IsDeleted         bool      `json:"is_deleted"`
	DeletedAt         time.Time `json:"deleted_at"`
	CreatedAt         time.Time `json:"created_at"`
}
