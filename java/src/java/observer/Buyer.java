package observer;

public abstract class Buyer {
    protected String name;
    protected Shop shop;

    public Buyer(String name, Shop shop) {
        this.name = name;
        this.shop = shop;
    }

    public abstract void inform();
}
