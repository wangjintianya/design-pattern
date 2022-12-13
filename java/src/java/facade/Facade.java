package facade;

public class Facade {
    private final Chef chef;
    private final Waiter waiter;
    private final Cleaner cleaner;

    public Facade() {
        VegVendor vv = new VegVendor();
        vv.sell();
        this.chef = new Chef();
        this.waiter = new Waiter();
        this.cleaner = new Cleaner();
    }

    public void provideService() {
        // 接待、点菜
        waiter.order();
        // 厨师做饭
        chef.cook();
        // 上菜
        waiter.serve();
        //收拾桌子，洗碗
        cleaner.clean();
        cleaner.wash();
    }
}
