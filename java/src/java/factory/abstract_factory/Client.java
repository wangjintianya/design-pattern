package factory.abstract_factory;

public class Client {
    public static void main(String[] args) {
        System.out.println("游戏开始。。。");
        System.out.println("双方挖矿攒钱。。。");
        //第一位玩家选择了地球人族
        System.out.println("工人建造人族工厂。。。");
        AbstractFactory factory = new HumanFactory(10, 10);
        Unit marine = factory.createLowClass();
        marine.show();
        Unit tank = factory.createMidClass();
        tank.show();
        Unit battleShip = factory.createHighClass();
        battleShip.show();

        //另一位玩家选择了外星族
        System.out.println("工蜂建造外星虫族工厂。。。");
        factory = new AlienFactory(200, 200);
        Unit roach = factory.createLowClass();
        roach.show();
        Unit spitter = factory.createMidClass();
        spitter.show();
        Unit mammoth = factory.createHighClass();
        mammoth.show();

        System.out.println("两族开始大混战。。。");
        marine.attack();
        roach.attack();
        spitter.attack();
        tank.attack();
        mammoth.attack();
        battleShip.attack();
    }
}
