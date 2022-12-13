package composite;

public abstract class Node {
    protected String name;

    public Node(String name) {
        this.name = name;
    }

    // 添加后续子节点的方法
    protected abstract void add(Node node);

    protected void ls(int space) {
        for (int i = 0; i < space; i++) {
            System.out.print(" ");
        }
        System.out.println(name);
    }
}
