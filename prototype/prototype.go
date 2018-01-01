package prototype

type producter interface {
	clone() producter
	GetName() string
}

type Manager struct {
	product producter
}

func (m *Manager) Register(producter producter) {
	m.product = producter
}

func (m *Manager) Create(name string) producter {
	producter := m.product
	return producter.clone()
}

type Product struct {
	name string
}

func (p *Product) SetUp() {
	// something takes time...
}

func (p *Product) GetName() string {
	return p.name
}

// 新しい構造体に自身の値をセットして返すことで擬似的にcloneとした。
// ポインタ参照まで考慮したdeepcopyに関しては実装もしくは機能を提供する
// パッケージが必要になる
func (p *Product) clone() producter {
	return &Product{p.name}
}
