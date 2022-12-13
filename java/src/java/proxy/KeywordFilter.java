package proxy;

import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.util.Arrays;
import java.util.List;

public class KeywordFilter implements InvocationHandler {

    private final List<String> blackList = Arrays.asList("电影", "游戏", "小说", "音乐");

    private final Object origin;

    public KeywordFilter(Object origin) {
        this.origin = origin;//注入被代理对象
        System.out.println("开启关键字过滤模式...");
    }

    @Override
    public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
        //要被切入方法面之前的业务逻辑
        String arg = args[0].toString();
        for (String keyword : blackList) {
            if (arg.contains(keyword)) {
                System.out.println("禁止访问：" + arg);
                return null;
            }
        }

        return method.invoke(origin, arg);
    }
}
