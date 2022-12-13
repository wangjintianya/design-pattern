package flyweight;

public class Client {
    public static void main(String[] args) {
        Factory factory = new Factory();
        factory.getDrawable("河流").draw(10, 10);
        factory.getDrawable("河流").draw(10, 20);
        factory.getDrawable("石路").draw(10, 30);
        factory.getDrawable("草坪").draw(10, 40);
        factory.getDrawable("草坪").draw(10, 50);
        factory.getDrawable("草坪").draw(10, 60);
        factory.getDrawable("草坪").draw(10, 70);
        factory.getDrawable("草坪").draw(10, 80);
        factory.getDrawable("房子").draw(10, 80);
    }
}
