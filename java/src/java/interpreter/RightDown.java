package interpreter;

public class RightDown implements Expression{
    @Override
    public void interpret() {
        System.out.println("按下鼠标：右键");
    }
}
