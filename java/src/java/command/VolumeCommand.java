package command;

public class VolumeCommand implements Command{
    private final Device device;

    public VolumeCommand(Device device) {
        this.device = device;
    }

    @Override
    public void exe() {
        device.volumeUp();
    }

    @Override
    public void unexe() {
        device.volumeDown();
    }
}
