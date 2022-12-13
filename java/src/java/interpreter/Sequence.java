package interpreter;

import java.util.List;

public class Sequence implements Expression{

    private final List<Expression> expressions;

    public Sequence(List<Expression> expressions) {
        this.expressions = expressions;
    }

    @Override
    public void interpret() {
        expressions.forEach(Expression::interpret);
    }
}
