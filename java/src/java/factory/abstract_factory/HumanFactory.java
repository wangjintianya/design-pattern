package factory.abstract_factory;

public class HumanFactory implements AbstractFactory {
    private int x, y;// 人族工厂坐标

    public HumanFactory(int x, int y) {
        this.x = x;
        this.y = y;
    }

    @Override
    public Unit createLowClass() {
        Unit unit = new Marine(x, y);
        System.out.println("制作海军陆战队员成功。");
        return unit;
    }

    @Override
    public Unit createMidClass() {
        Unit unit = new Tank(x, y);
        System.out.println("制造变形坦克成功。");
        return unit;
    }

    @Override
    public Unit createHighClass() {
        Unit unit = new BattleShip(x, y);
        System.out.println("制造巨型战舰成功。");
        return unit;
    }
}
