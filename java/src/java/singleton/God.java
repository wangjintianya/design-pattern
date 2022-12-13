package singleton;

public class God {

    private static God god;

    private God() {
    }

    public static God GetInstance() {
        if (god == null) {
            synchronized(God.class){
                if (god == null) {//只有头香造了神，其他抢香的白排队了
                    god = new God();
                }
            }
        }
        return god;
    }
}
