package interpreter;

public class Delay implements Expression{

    private int seconds;

    public Delay(int seconds) {
        this.seconds = seconds;
    }

    @Override
    public void interpret() {
        System.out.println("系统延时：" + seconds + "秒");
        try {
            Thread.sleep(seconds * 1000L);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}
