package state;

public class Client {
    public static void main(String[] args) {
        Switcher switcher = new Switcher();
        switcher.switchOn();
        switcher.switchOn();
        switcher.switchOff();
        switcher.switchOff();
    }
}
