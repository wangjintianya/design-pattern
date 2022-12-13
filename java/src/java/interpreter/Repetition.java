package interpreter;

public class Repetition implements Expression{

    private int loopCount;
    private Expression expression;

    public Repetition(int loopCount, Expression expression) {
        this.loopCount = loopCount;
        this.expression = expression;
    }

    @Override
    public void interpret() {
        while (loopCount > 0) {
            expression.interpret();
            loopCount--;
        }
    }
}
