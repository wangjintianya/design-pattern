package command;

public class Client {
    public static void main(String[] args) {
        TV tv = new TV();
        Controller controller = new Controller();
        controller.setOkCommand(new SwitchCommand(tv));
        controller.setVerticalCommand(new VolumeCommand(tv));
        controller.setHorizontalCommand(new ChannelCommand(tv));
        controller.buttonOKHold();
        controller.buttonOKClick();
        controller.buttonLeftClick();
        controller.buttonLeftClick();
        controller.buttonRightClick();
        controller.buttonRightClick();
        controller.buttonUpClick();
        controller.buttonUpClick();
        controller.buttonDownClick();
        controller.buttonDownClick();

        System.out.println("----------------------------");
        Radio radio = new Radio();
        controller.setOkCommand(new SwitchCommand(radio));
        controller.setVerticalCommand(new VolumeCommand(radio));
        controller.setHorizontalCommand(new ChannelCommand(radio));

        controller.buttonOKHold();
        controller.buttonOKClick();
        controller.buttonLeftClick();
        controller.buttonRightClick();
        controller.buttonUpClick();
        controller.buttonDownClick();
    }
}
