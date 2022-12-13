package memento;

import java.util.ArrayList;
import java.util.List;

public class Editor {
    private final Doc doc;
    private final List<History> histories;
    private int historyPosition = -1;

    public Editor(Doc doc) {
        System.out.println("<<<打开文档" + doc.getTitle());
        histories = new ArrayList<>();
        this.doc = doc;
        backup();
        show();
    }

    public void backup() {
        histories.add(doc.createHistory());
        historyPosition++;
    }

    private void show() {
        System.out.println(doc.getBody());
        System.out.println("文章结束>>>");
    }

    public void append(String txt) {
        System.out.println("<<<插入操作");
        backup();
        doc.setBody(doc.getBody() + txt);
        show();
    }

    public void save() {
        System.out.println("<<<存盘操作");
    }

    public void delete() {
        System.out.println("<<<删除操作");
        backup();
        doc.setBody("");
        show();
    }

    public void undo() {
        System.out.println(">>>撤销操作");
        if (historyPosition == 0) {
            return;
        }
        History history = histories.get(historyPosition);
        historyPosition--;
        doc.restoreHistory(history);
        show();
    }
}
