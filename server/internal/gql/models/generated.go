// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type NewProduct struct {
	Sku            string  `json:"SKU"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Price          float64 `json:"price"`
	Stock          int     `json:"stock"`
	StartDate      string  `json:"startDate"`
	EndDate        string  `json:"endDate"`
	CreatedByAdmin int     `json:"createdByAdmin"`
}

type NewUser struct {
	Email        string  `json:"email"`
	Token        string  `json:"token"`
	AuthMethodID int     `json:"authMethodId"`
	FirstName    *string `json:"firstName"`
	LastName     *string `json:"lastName"`
}

type Product struct {
	Sku            string  `json:"SKU"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Price          float64 `json:"price"`
	Stock          int     `json:"stock"`
	StartDate      string  `json:"startDate"`
	EndDate        string  `json:"endDate"`
	CreatedByAdmin *User   `json:"createdByAdmin"`
	CreatedAt      *string `json:"createdAt"`
	UpdatedAt      *string `json:"updatedAt"`
}

type UpdatedProfile struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type User struct {
	Email       string       `json:"email"`
	Role        int          `json:"role"`
	CreatedAt   *string      `json:"createdAt"`
	UpdatedAt   *string      `json:"updatedAt"`
	UserProfile *UserProfile `json:"userProfile"`
}

type UserProfile struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Email     *string `json:"email"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt *string `json:"updatedAt"`
}
