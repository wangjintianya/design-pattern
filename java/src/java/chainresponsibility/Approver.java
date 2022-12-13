package chainresponsibility;

public abstract class Approver {
    protected String name;
    protected Approver nextApprover;

    public Approver(String name) {
        this.name = name;
    }

    protected  Approver setNextApprover(Approver nextApprover) {
        this.nextApprover = nextApprover;
        return nextApprover;
    }

    public abstract void approve(int amount);
}
