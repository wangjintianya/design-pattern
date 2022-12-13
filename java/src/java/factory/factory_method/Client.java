package factory.factory_method;

public class Client {
    public static void main(String[] args) {
        SimpleFactory factory = new SimpleFactory(800);
        Enemy enemy = factory.create("Tank");
        Enemy enemy1 = factory.create("Tank");
        Enemy enemy2 = factory.create("AirPlane");
        enemy.show();
        enemy1.show();
        enemy2.show();
    }
}
