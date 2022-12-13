package chain_responsibility

import "fmt"

type approve interface {
	HasRight(amount int) bool
	ApprovePass(amount int)
	ApproveFail(amount int)
}

type Approver struct {
	approve
	nextApprover *Approver
}

func (a *Approver) Approve(amount int) {
	if a.approve.HasRight(amount) {
		a.approve.ApprovePass(amount)
	} else {
		a.approve.ApproveFail(amount)
		if a.nextApprover != nil {
			a.nextApprover.Approve(amount)
		}
	}
}

type CEO struct {
	name string
}

func (c *CEO) HasRight(amount int) bool {
	if amount < 10000 {
		return true
	}
	return false
}

func (c *CEO) ApprovePass(amount int) {
	fmt.Printf("审批通过.【CEO：%s】\n", c.name)
}

func (c *CEO) ApproveFail(amount int) {
	fmt.Printf("审批不通过.【CEO：%s】\n", c.name)
}

func NewCEO(name string) *Approver {
	return &Approver{
		approve: &CEO{
			name: name,
		},
	}
}

type Manager struct {
	name string
}

func (m *Manager) HasRight(amount int) bool {
	if amount < 5000 {
		return true
	}
	return false
}

func (m *Manager) ApprovePass(amount int) {
	fmt.Printf("审批通过.【经理：%s】\n", m.name)
}

func (m *Manager) ApproveFail(amount int) {
	fmt.Printf("无权审批，升级处理。【经理：%s】\n", m.name)
}

func NewManager(name string) *Approver {
	return &Approver{
		approve: &Manager{name: name},
	}
}

type Staff struct {
	name string
}

func (s *Staff) HasRight(amount int) bool {
	if amount < 1000 {
		return true
	}
	return false
}

func (s *Staff) ApprovePass(amount int) {
	fmt.Printf("审批通过.【员工：%s】\n", s.name)
}

func (s *Staff) ApproveFail(amount int) {
	fmt.Printf("无权审批，升级处理。【员工：%s】\n", s.name)
}

func NewStaff(name string) *Approver {
	return &Approver{
		approve: &Staff{name: name},
	}
}
