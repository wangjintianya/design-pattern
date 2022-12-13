package visitor;

import java.time.LocalDate;

public abstract class Product {
    protected String name;
    protected LocalDate productDate;
    protected float price;

    public Product(String name, LocalDate productDate, float price) {
        this.name = name;
        this.productDate = productDate;
        this.price = price;
    }

    public String getName() {
        return name;
    }

    public LocalDate getProductDate() {
        return productDate;
    }

    public float getPrice() {
        return price;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void setProductDate(LocalDate productDate) {
        this.productDate = productDate;
    }

    public void setPrice(float price) {
        this.price = price;
    }
}
