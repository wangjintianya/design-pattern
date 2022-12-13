package visitor;

import java.time.LocalDate;

public class Wine extends Product implements Acceptable{
    @Override
    public void accept(Visitor visitor) {
        visitor.visit(this);
    }

    public Wine(String name, LocalDate productDate, float price) {
        super(name, productDate, price);
    }
}
