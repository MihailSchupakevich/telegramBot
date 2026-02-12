package handlers

type Product struct {
	nameProduct string
	count       int
}

type ProductsI interface {
	addProduct()
	stopAddProduct()
	listProducts()
}
