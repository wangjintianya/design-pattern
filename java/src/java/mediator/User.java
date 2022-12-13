package mediator;

public class User {
    private final String name;
    private ChatRoom chatRoom;

    public String getName() {
        return name;
    }

    public User(String name) {
        this.name = name;
    }

    public void login(ChatRoom chatRoom) {
        chatRoom.connect(this);
        this.chatRoom = chatRoom;
    }

    public void talk(String message) {
        chatRoom.sendMessage(this, message);
    }

    public void listen(User user, String message) {
        System.out.print("【" + this.name + "的对话框】");
        System.out.println(user.getName() + " 说： " + message);
    }
}
