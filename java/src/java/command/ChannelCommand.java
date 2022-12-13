package command;

public class ChannelCommand implements Command{
    private final Device device;

    public ChannelCommand(Device device) {
        this.device = device;
    }

    @Override
    public void exe() {
        device.channelUp();
    }

    @Override
    public void unexe() {
        device.channelDown();
    }
}
