package factory.abstract_factory;

public class AlienFactory implements AbstractFactory {
    private int x, y;// 异族工厂坐标

    public AlienFactory(int x, int y) {
        this.x = x;
        this.y = y;
    }

    @Override
    public Unit createLowClass() {
        Unit unit = new Roach(x, y);
        System.out.println("制作蟑螂兵成功。");
        return unit;
    }

    @Override
    public Unit createMidClass() {
        Unit unit = new Spitter(x, y);
        System.out.println("制造毒液兵成功。");
        return unit;
    }

    @Override
    public Unit createHighClass() {
        Unit unit = new Mammoth(x, y);
        System.out.println("制造猛犸巨兽成功。");
        return unit;
    }
}
