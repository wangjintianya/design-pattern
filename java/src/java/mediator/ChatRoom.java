package mediator;

import java.util.ArrayList;
import java.util.List;

public class ChatRoom {
    private String name;

    public ChatRoom(String name) {
        this.name = name;
    }

    public List<User> users = new ArrayList<>();//聊天室里的用户们

    public void connect(User user) {
        users.add(user);
        System.out.print("欢迎【");
        System.out.print(user.getName());
        System.out.println("】加入聊天室【" + name + "】");
    }

    public void sendMessage(User user, String message) {
        users.stream().filter(user1 -> !user.equals(user1)).forEach(user1 -> user1.listen(user, message));
    }
}
