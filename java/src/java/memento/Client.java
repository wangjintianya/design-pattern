package memento;

public class Client {
    public static void main(String[] args) {
        Doc doc = new Doc("《AI的觉醒》");
        Editor editor = new Editor(doc);
        editor.append("第一章 混沌初开");
        editor.append("\n  正文2000字……");
        editor.append("\n第二章 荒漠之花\n  正文3000字……");

        editor.undo();
        editor.undo();
        editor.undo();
    }
}
