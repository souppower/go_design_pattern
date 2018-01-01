package factoryMethod

type creater interface {
	createProduct(owner string) User
	registerProduct(User)
}

// User interface
type User interface {
	Use() string
}

// Factory struct
type Factory struct {
}

// インスタンス生成部分にはTemplateMethodパターンを使う。
// Goではclient-specified self patternにより親構造体のメソッド内で子構造体の
// 実装を呼ぶことができる
//
// 処理の流れが簡易でテンプレートを使わないでよければ、ファクトリとなる関数を
// 直接引数に渡してもよい
// この場合、個別のファクトリークラスを定義しなくてもよくなる。
// func (self *Factory) Create(factoryMethod func() User) {
//   self.product = factoryMethod()
//   return self.product
// }

// Create returns a User instance
func (f *Factory) Create(factory creater, owner string) User {
	user := factory.createProduct(owner)
	factory.registerProduct(user)
	return user
}

// IDCard struct
type IDCard struct {
	owner string
}

// Use returns the IDCard's owner
func (ic *IDCard) Use() string {
	return ic.owner
}

// IDCardFactory is a concrete implementation of Factory
type IDCardFactory struct {
	*Factory
	owners []string
}

func (icf *IDCardFactory) createProduct(owner string) User {
	return &IDCard{owner}
}

func (icf *IDCardFactory) registerProduct(product User) {
	owner := product.Use()
	icf.owners = append(icf.owners, owner)
}
