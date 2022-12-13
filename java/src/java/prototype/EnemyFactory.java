package prototype;

public class EnemyFactory {
    private static final EnemyPlane protoType = new EnemyPlane(200);

    // 获取敌机克隆实例
    public static EnemyPlane getInstance(int x) throws CloneNotSupportedException {
        EnemyPlane clone = protoType.clone();
        clone.setX(x);
        return clone;
    }
}
