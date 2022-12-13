package factory.abstract_factory;

public class Marine extends Unit{ // 海军陆战队士兵
    public Marine(int x, int y) {
        super(6, 5, 40, x, y);
    }

    @Override
    public void show() {
        System.out.println("士兵出现在坐标：[" + x + "," + y + "]");
    }

    @Override
    public void attack() {
        System.out.println("士兵用机关枪射击，攻击力：" + attack);
    }
}
