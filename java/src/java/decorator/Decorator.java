package decorator;

public abstract class Decorator implements Showable {
    public Showable showable;

    public Decorator(Showable showable) {
        this.showable = showable;
    }

    @Override
    public void show() {
        showable.show(); //这家伙素面朝天的秀
    }
}
