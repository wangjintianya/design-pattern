package prototype;

public class EnemyPlane implements Cloneable {
    private int x;
    private int y = 0;

    public EnemyPlane(int x) {
        this.x = x;
    }

    public int getX() {
        return x;
    }

    public int getY() {
        return y;
    }

    public void setX(int x) {
        this.x = x;
    }

    public void fly () {
        y++;
    }

    @Override
    protected EnemyPlane clone() throws CloneNotSupportedException {
        return (EnemyPlane)super.clone();
    }
}
