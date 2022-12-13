package prototype

//Cloneable 是原型对象需要实现的接口
type Cloneable interface {
	Clone() Cloneable
}

type Manager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *Manager {
	return &Manager{
		prototypes: make(map[string]Cloneable),
	}
}

func (p *Manager) Get(name string) Cloneable {
	return p.prototypes[name].Clone()
}

func (p *Manager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}

type EnemyPlane struct {
	X, Y int
}

func (e *EnemyPlane) Clone() Cloneable {
	newKeyword := *e
	return &newKeyword
}
