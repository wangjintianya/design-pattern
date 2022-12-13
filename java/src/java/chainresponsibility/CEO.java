package chainresponsibility;

public class CEO extends Approver {
    public CEO(String name) {
        super(name);
    }

    @Override
    public void approve(int amount) {
        if (amount < 10000) {
            System.out.println("审批通过。【CEO：" + name + "】");
        } else {
            System.out.println("审批不通过。【CEO：" + name + "】");
        }
    }
}
