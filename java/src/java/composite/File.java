package composite;

public class File extends Node{
    public File(String name) {
        super(name);
    }

    @Override
    public void add(Node node) {
        System.out.println("文件不可添加子节点");
    }

    @Override
    public void ls(int space) {
        super.ls(space);
    }

}
