package interpreter;

import java.util.Arrays;

public class Client {
    public static void main(String[] args) {
        new Move(500, 600);
        new Repetition(5, new Sequence(Arrays.asList(
                new LeftClick(),
                new Delay(1)
        ))).interpret();
        new RightDown().interpret();
        new Delay(7).interpret();
    }
}
