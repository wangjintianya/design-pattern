package factory.abstract_factory;

public class Spitter extends Unit{ // 外星毒液口水兵
    public Spitter(int x, int y) {
        super(10, 8, 80, x, y);
    }

    @Override
    public void show() {
        System.out.println("口水兵出现在坐标：[" + x + "," + y + "]");
    }

    @Override
    public void attack() {
        System.out.println("口水兵用毒液喷射，攻击力：" + attack);
    }
}
