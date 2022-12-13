package visitor;

import java.time.LocalDate;

public class Fruit extends Product implements Acceptable{

    private float weight;

    public Fruit(String name, LocalDate productDate, float price, float weight) {
        super(name, productDate, price);
        this.weight = weight;
    }

    public void setWeight(float weight) {
        this.weight = weight;
    }

    public float getWeight() {
        return weight;
    }

    @Override
    public void accept(Visitor visitor) {
        visitor.visit(this);
    }
}
