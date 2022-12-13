package factory.factory_method;

import java.util.Random;

public class SimpleFactory {
    private final int screenWith;
    private final Random random;

    public SimpleFactory(int screenWith) {
        this.screenWith = screenWith;
        this.random = new Random();
    }

    public Enemy create(String type) {
        int x = random.nextInt(screenWith);
        Enemy enemy = null;
        switch (type) {
            case "AirPlane":
                enemy = new AirPlane(x, 0);
                break;
            case "Tank":
                enemy = new Tank(x, 0);

        }
        return enemy;
    }
}