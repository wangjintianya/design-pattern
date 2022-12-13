package bridge;

public class Client {
    public static void main(String[] args) {
        // 黑色画笔对应的所有形状
        new BlackPen(new CircleRuler()).draw();
        new BlackPen(new SquareRuler()).draw();
        new BlackPen(new TriangleRuler()).draw();
        // 红色画笔对应的所有形状
        new RedPen(new CircleRuler()).draw();
        new RedPen(new SquareRuler()).draw();
        new RedPen(new TriangleRuler()).draw();
    }
}
