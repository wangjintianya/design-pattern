package flyweight;

public class Grass implements Drawable{

    private final String image;

    public Grass() {
        this.image = "草坪";
        System.out.println("从磁盘加载[" + this.image + "]图片，耗时半秒。。。");
    }

    @Override
    public void draw(int x, int y) {
        System.out.println("在位置[" + x + "," + y + "]上绘制图片：[" + image + "]");
    }
}
