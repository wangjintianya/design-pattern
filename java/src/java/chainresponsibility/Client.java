package chainresponsibility;

public class Client {
    public static void main(String[] args) {
        Staff staff = new Staff("张飞");
        staff.setNextApprover(new Manager("关羽")).setNextApprover(new CEO("刘备"));
        staff.approve(100);
        staff.approve(2000);
        staff.approve(5000);
        staff.approve(10000);
    }
}
