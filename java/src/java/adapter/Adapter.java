package adapter;

public class Adapter implements TriplePin{

    private final DualPin dualPinDevice;

    public Adapter(DualPin dualPinDevice) {
        this.dualPinDevice = dualPinDevice;
    }

    @Override
    public void electrify(int l, int n, int e) {
        dualPinDevice.electrify(l, n);
    }
}
