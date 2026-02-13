package handlers

type Product struct {
	nameProduct string
	count       string
}

type ProductsI interface {
	addProduct()
	stopAddProduct()
	listProducts()
}
