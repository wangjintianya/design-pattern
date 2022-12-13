package factory.abstract_factory;

public interface AbstractFactory {
    public Unit createLowClass();// 工厂方法：制造低级兵种

    public Unit createMidClass();// 工厂方法：制造中级兵种

    public Unit createHighClass();// 工厂方法：制造高级兵种
}
