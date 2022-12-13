package factory.abstract_factory;

public class BattleShip extends Unit{ // 巨型战舰
    public BattleShip(int x, int y) {
        super(35, 200, 500, x, y);
    }

    @Override
    public void show() {
        System.out.println("战舰出现在坐标：[" + x + "," + y + "]");
    }

    @Override
    public void attack() {
        System.out.println("战舰用激光炮打击，攻击力：" + attack);
    }
}
