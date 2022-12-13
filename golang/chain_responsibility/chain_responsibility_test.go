package chain_responsibility

import "testing"

func TestChainResponsibility(t *testing.T) {
	ceo := NewCEO("刘备")
	manager := NewManager("关羽")
	staff := NewStaff("张飞")
	manager.nextApprover = ceo
	staff.nextApprover = manager
	staff.Approve(100)
	staff.Approve(1000)
	staff.Approve(100000)
}
