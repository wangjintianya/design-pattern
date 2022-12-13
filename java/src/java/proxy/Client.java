package proxy;

public class Client {
    public static void main(String[] args) {
        RouterProxy routerProxy = new RouterProxy();
        routerProxy.access("http://www.电影.com");
        routerProxy.access("http://www.小说.com");
        routerProxy.access("http://www.读书.com");
        routerProxy.access("http://www.游戏.com");
    }
}
