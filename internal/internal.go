package internal

// Calculator struct untuk menyimpan nilai dasar
type Calculator struct {
	Base float32
}

// NewCalculator menginisialisasi Calculator dengan nilai awal
func NewCalculator(base float32) *Calculator {
	return &Calculator{Base: base}
}

// Add menambahkan angka ke Base
func (c *Calculator) Add(num float32) {
	c.Base += num
}

// Subtract mengurangi angka dari Base
func (c *Calculator) Subtract(num float32) {
	c.Base -= num
}

// Multiply mengalikan Base dengan angka
func (c *Calculator) Multiply(num float32) {
	c.Base *= num
}

// Divide membagi Base dengan angka, menghindari pembagian dengan nol
func (c *Calculator) Divide(num float32) {
	if num != 0 {
		c.Base /= num
	}
}

// Result mengembalikan nilai saat ini dari Base
func (c *Calculator) Result() float32 {
	return c.Base
}
