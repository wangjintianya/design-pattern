package observer;

import java.util.ArrayList;
import java.util.List;

public class Shop {
    private String product;
    private List<Buyer> buyers;

    public Shop() {
        this.product = "无产品";
        this.buyers = new ArrayList<>();
    }

    // 注册
    public void register(Buyer buyer) {
        this.buyers.add(buyer);
    }

    public String getProduct() {
        return product;
    }

    public void setProduct(String product) {
        this.product = product;
        notifyBuyers();
    }

    public void notifyBuyers() {
        buyers.forEach(Buyer::inform);
    }
}
