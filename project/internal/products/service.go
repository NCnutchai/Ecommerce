package products

type Service interface {
	ProductList() (*[]Product, error)
	ProductDetail(code string) (*Product, error)
	ProductCreate(product Product) error
}

type productService struct {
	repo Repository
}

func NewProductService(repo Repository) Service {
	return &productService{
		repo: repo,
	}
}

func (s *productService) ProductList() (*[]Product, error) {
	return s.repo.FindAll()
}

func (s *productService) ProductDetail(code string) (*Product, error) {
	return s.repo.FindByCode(code)
}

func (s *productService) ProductCreate(product Product) error {
	return s.repo.Create(product)
}
