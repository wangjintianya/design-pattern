package flyweight;

public class Stone implements Drawable{

    private final String image;

    public Stone() {
        this.image = "石头";
        System.out.println("从磁盘加载[" + this.image + "]图片，耗时半秒。。。");
    }

    @Override
    public void draw(int x, int y) {
        System.out.println("在位置[" + x + "," + y + "]上绘制图片：[" + image + "]");
    }
}
