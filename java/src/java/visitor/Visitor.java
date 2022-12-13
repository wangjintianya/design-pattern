package visitor;

public interface Visitor {
    public void visit(Candy candy);
    public void visit(Fruit fruit);
    public void visit(Wine wine);
}
