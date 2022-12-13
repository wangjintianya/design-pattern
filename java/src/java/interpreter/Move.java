package interpreter;

public class Move implements Expression{
    private final int x;
    private final int y;

    public Move(int x, int y) {
        this.x = x;
        this.y = y;
    }

    @Override
    public void interpret() {
        System.out.println("移动鼠标：【" + x + "," + y +"】");
    }
}
