package visitor;

import java.time.LocalDate;

public class Candy extends Product implements Acceptable{
    public Candy(String name, LocalDate productDate, float price) {
        super(name, productDate, price);
    }

    @Override
    public void accept(Visitor visitor) {
        visitor.visit(this);
    }
}
