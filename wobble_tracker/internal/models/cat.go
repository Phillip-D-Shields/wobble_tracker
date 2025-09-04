package models

import "time"

type Cat struct {
    ID            int       `json:"id" db:"id"`
    Name          string    `json:"name" db:"name"`
    Breed         *string   `json:"breed" db:"breed"`
    Color         *string   `json:"color" db:"color"`
    Age           *int      `json:"age" db:"age"`
    Weight        *float64  `json:"weight" db:"weight"`
    IsIndoor      bool      `json:"is_indoor" db:"is_indoor"`
    IsVaccinated  bool      `json:"is_vaccinated" db:"is_vaccinated"`
    MicrochipID   *string   `json:"microchip_id" db:"microchip_id"`
    OwnerName     *string   `json:"owner_name" db:"owner_name"`
    OwnerEmail    *string   `json:"owner_email" db:"owner_email"`
    Notes         *string   `json:"notes" db:"notes"`
    CreatedAt     time.Time `json:"created_at" db:"created_at"`
    UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

// payload for creating a new cat
type CreateCatRequest struct {
    Name         string   `json:"name"`
    Breed        *string  `json:"breed,omitempty"`
    Color        *string  `json:"color,omitempty"`
    Age          *int     `json:"age,omitempty"`
    Weight       *float64 `json:"weight,omitempty"`
    IsIndoor     *bool    `json:"is_indoor,omitempty"`
    IsVaccinated *bool    `json:"is_vaccinated,omitempty"`
    MicrochipID  *string  `json:"microchip_id,omitempty"`
    OwnerName    *string  `json:"owner_name,omitempty"`
    OwnerEmail   *string  `json:"owner_email,omitempty"`
    Notes        *string  `json:"notes,omitempty"`
}

// payload for updating a cat
type UpdateCatRequest struct {
    Name         *string  `json:"name,omitempty"`
    Breed        *string  `json:"breed,omitempty"`
    Color        *string  `json:"color,omitempty"`
    Age          *int     `json:"age,omitempty"`
    Weight       *float64 `json:"weight,omitempty"`
    IsIndoor     *bool    `json:"is_indoor,omitempty"`
    IsVaccinated *bool    `json:"is_vaccinated,omitempty"`
    MicrochipID  *string  `json:"microchip_id,omitempty"`
    OwnerName    *string  `json:"owner_name,omitempty"`
    OwnerEmail   *string  `json:"owner_email,omitempty"`
    Notes        *string  `json:"notes,omitempty"`
}

// filtering options for cat queries
type CatFilter struct {
    Breed        *string `json:"breed,omitempty"`
    IsIndoor     *bool   `json:"is_indoor,omitempty"`
    IsVaccinated *bool   `json:"is_vaccinated,omitempty"`
    MinAge       *int    `json:"min_age,omitempty"`
    MaxAge       *int    `json:"max_age,omitempty"`
    OwnerEmail   *string `json:"owner_email,omitempty"`
}

// pagination and filtering for cat listing
type ListCatsRequest struct {
    Limit  int        `json:"limit"`
    Offset int        `json:"offset"`
    Filter *CatFilter `json:"filter,omitempty"`
}