package adapter;

public class Client {
    public static void main(String[] args) {
        TV tv = new TV();
        Adapter adapter = new Adapter(tv);
        adapter.electrify(1,2,3);
    }
}
