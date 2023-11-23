package product

import "time"

type Product struct {
	ID          uint64     `db:"id" json:"id"`
	Title       string     `db:"title" json:"title"`
	Sku         string     `db:"sku" json:"sku"`
	Description string     `db:"description" json:"description"`
	Category    string     `db:"category" json:"category"`
	Weight      float64    `db:"weight" json:"weight"`
	Price       float64    `db:"price" json:"price"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
}

type ProductImage struct {
	ID               uint64 `db:"id" json:"id"`
	ProductID        uint64 `db:"product_id" json:"product_id"`
	ImageUrl         string `db:"image_url" json:"image_url"`
	ImageDescription string `db:"image_description" json:"image_description"`
}

type ProductReview struct {
	ID        uint64 `db:"id" json:"id"`
	ProductID uint64 `db:"product_id" json:"product_id"`
	Rating    uint8  `db:"rating" json:"rating"`
	Comment   string `db:"comment" json:"comment"`
}

type ProductDetail struct {
	Product
	Images  *[]ProductImage  `json:"images,omitempty"`
	Reviews *[]ProductReview `json:"reviews,omitempty"`
	Rating  float64          `db:"rating" json:"rating"`
}
