package proxy;

public class Modem implements Internet{
    @Override
    public void access(String url) {
        System.out.println("正在访问:" + url);
    }
}
