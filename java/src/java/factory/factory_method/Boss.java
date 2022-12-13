package factory.factory_method;

public class Boss extends Enemy{
    public Boss(int x, int y) {
        super(x, y);
    }

    @Override
    public void show() {
        System.out.println("Boss出现坐标：" + x + "," + y);
        System.out.println("Boss向玩家发起攻击……");
    }
}
