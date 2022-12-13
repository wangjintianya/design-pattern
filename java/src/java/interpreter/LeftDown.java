package interpreter;

public class LeftDown implements Expression{
    @Override
    public void interpret() {
        System.out.println("按下鼠标：左键");
    }
}
