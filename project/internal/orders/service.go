package orders

type Service interface {
	CreateOrder(order Order) (Order, error)
	CancelOrder(order_number string, user_id int) error
	OrderDetail(order_number string, user_id int) (Order, error)
	GetAllOrders() ([]Order, error)
	GetOrderHistory(user_id int) ([]Order, error)
}

type orderService struct {
	repo Repository
}

func NewOrderService(repo Repository) Service {
	return &orderService{
		repo: repo,
	}
}

func (s *orderService) CreateOrder(order Order) (Order, error) {
	err := s.repo.Create(order)
	if err != nil {
		return Order{}, err
	}

	return order, nil
}

func (s *orderService) CancelOrder(order_number string, user_id int) error {
	err := s.repo.Cancel(order_number, user_id)
	if err != nil {
		return err
	}

	return nil
}

func (s *orderService) OrderDetail(order_number string, user_id int) (Order, error) {
	return s.repo.Detail(order_number, user_id)
}

func (s *orderService) GetAllOrders() ([]Order, error) {
	return s.repo.FindAll()
}

func (s *orderService) GetOrderHistory(user_id int) ([]Order, error) {
	return s.repo.FindOrderByUserID(user_id)
}
