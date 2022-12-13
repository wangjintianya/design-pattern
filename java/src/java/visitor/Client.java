package visitor;

import java.time.LocalDate;
import java.util.Arrays;
import java.util.List;

public class Client {
    public static void main(String[] args) {
        List<Acceptable> lists = Arrays.asList(
                new Candy("小黑兔糖果", LocalDate.of(2022, 7, 1),1.1f),
                new Wine("猫泰白酒", LocalDate.of(2020, 3, 6),300f),
                new Fruit("草莓", LocalDate.of(2022, 10, 31),23f, 3.0f)
        );
        Visitor discountVisitor = new DiscountVisitor(LocalDate.of(2022, 11, 7));
        for (Acceptable acceptable : lists) {
            acceptable.accept(discountVisitor);
        }
    }
}
