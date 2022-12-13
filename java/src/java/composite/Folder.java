package composite;

import java.util.ArrayList;
import java.util.List;

public class Folder extends Node{

    private final List<Node> childrenNodes = new ArrayList<>();

    public Folder(String name) {
        super(name);
    }

    @Override
    public void add(Node node) {
        childrenNodes.add(node);
    }

    @Override
    public void ls(int space) {
        super.ls(space);
        space++;
        for (Node node : childrenNodes) {
            node.ls(space);
        }
    }
}
