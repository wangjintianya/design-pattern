package factory.abstract_factory;

public class Mammoth extends Unit{ // 外星猛犸巨兽兵
    public Mammoth(int x, int y) {
        super(10, 8, 80, x, y);
    }

    @Override
    public void show() {
        System.out.println("猛犸巨兽兵出现在坐标：[" + x + "," + y + "]");
    }

    @Override
    public void attack() {
        System.out.println("猛犸巨兽用獠牙顶，攻击力：" + attack);
    }
}
