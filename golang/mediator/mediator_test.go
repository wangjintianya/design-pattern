package mediator

import "testing"

func TestMediator(t *testing.T) {
	room := &ChatRoom{name: "~美少女~"}
	user := User{name: "小红"}
	user1 := User{name: "小兰"}
	user2 := User{name: "小白"}

	user.login(room)
	user1.login(room)
	user2.login(room)

	user.talk("在吗")
	user1.talk("在")
}
