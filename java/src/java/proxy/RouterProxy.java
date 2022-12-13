package proxy;

import java.util.Arrays;
import java.util.List;

public class RouterProxy implements Internet{
    private Modem modem;

    private List<String> blackLists = Arrays.asList("电影", "音乐", "游戏", "小说");

    public RouterProxy() {
        this.modem = new Modem();
        System.out.println("连接拨号上网成功。。。");
    }

    @Override
    public void access(String url) {
        for (String black: blackLists) {
            if (url.contains(black)) {
                System.out.println("禁止访问：" + url);
                return;
            }
        }
        modem.access(url);
    }
}
