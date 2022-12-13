package proxy;

import java.lang.reflect.Proxy;

public class DynamicClient {
    public static void main(String[] args) {
        // 访问外网，生成猫代理
        Internet internet = (Internet) Proxy.newProxyInstance(
                Modem.class.getClassLoader(),
                Modem.class.getInterfaces(),
                new KeywordFilter(new Modem())
        );
        internet.access("http://www.电影.com");
        internet.access("http://www.游戏.com");
        internet.access("http://www.学习.com");
        internet.access("http://www.工作.com");

        // 访问内网，生成交换机代理
        Intranet intranet = (Intranet) Proxy.newProxyInstance(
                Switch.class.getClassLoader(),
                Switch.class.getInterfaces(),
                new KeywordFilter(new Switch())
        );
        intranet.fileAccess("\\\\192.68.1.2\\共享\\电影\\IronHuman.mp4");
        intranet.fileAccess("\\\\192.68.1.2\\共享\\游戏\\Hero.exe");
        intranet.fileAccess("\\\\192.68.1.4\\shared\\Java学习资料.zip");
        intranet.fileAccess("\\\\192.68.1.6\\Java知音\\设计模式是什么鬼.doc");
    }
}
