package decorator;

public class Client {
    public static void main(String[] args) {
        Decorator decorator = new Lipstick(new FoundationMakeup(new Girl()));
        decorator.show();
    }
}
