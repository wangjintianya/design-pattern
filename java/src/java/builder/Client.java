package builder;

public class Client {
    public static void main(String[] args) {
        // 招工，建别墅
        Builder houseBuilder = new HouseBuilder();
        Director director = new Director(houseBuilder);
        System.out.println(director.direct());
        // 替换施工方，建公寓
        director.setBuilder(new ApartmentBuilder());
        System.out.println(director.direct());
    }
}
