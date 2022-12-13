package factory.factory_method;

public class AirPlane extends Enemy{
    public AirPlane(int x, int y) {
        super(x, y);
    }

    @Override
    public void show() {
        System.out.println("飞机出现的坐标：" + x + "," + y);
        System.out.println("飞机向玩家发起攻击。。。");
    }
}
