package interpreter;

public class LeftClick implements Expression{

    private LeftUp leftUp;
    private LeftDown leftDown;

    public LeftClick() {
        this.leftUp = new LeftUp();
        this.leftDown = new LeftDown();
    }

    @Override
    public void interpret() {
        leftDown.interpret();
        leftUp.interpret();
    }
}
