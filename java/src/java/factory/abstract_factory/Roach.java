package factory.abstract_factory;

public class Roach extends Unit{ // 外星蟑螂兵
    public Roach(int x, int y) {
        super(5, 2, 35, x, y);
    }

    @Override
    public void show() {
        System.out.println("蟑螂兵出现在坐标：[" + x + "," + y + "]");
    }

    @Override
    public void attack() {
        System.out.println("蟑螂兵用爪子挠，攻击力：" + attack);
    }
}
