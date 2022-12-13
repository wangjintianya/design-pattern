package mediator

import "fmt"

type ChatRoom struct {
	name  string
	users []*User
}

func (c *ChatRoom) connect(user *User) {
	fmt.Printf("欢迎【%s】加入聊天室【%s】\n", user.name, c.name)
	c.users = append(c.users, user)
}

func (c *ChatRoom) sendMessage(fromUser *User, message string) {
	for _, user := range c.users {
		if user != fromUser {
			user.listen(fromUser, message)
		}
	}
}

type User struct {
	name     string
	chatRoom *ChatRoom
}

func (u *User) login(room *ChatRoom) {
	room.connect(u)
	u.chatRoom = room
}

func (u *User) talk(message string) {
	u.chatRoom.sendMessage(u, message)
}

func (u *User) listen(user *User, message string) {
	fmt.Printf("【%s 的对话框】 %s 说：%s\n", u.name, user.name, message)
}
