package mediator;

public class Client {
    public static void main(String[] args) {
        ChatRoom chatRoom = new ChatRoom("~美少女~");
        User user = new User("小红");
        User user1 = new User("小兰");
        User user2 = new User("小白");

        user.login(chatRoom);
        user1.login(chatRoom);
        user2.login(chatRoom);

        user.talk("在吗");
        user1.talk("在");
    }
}
