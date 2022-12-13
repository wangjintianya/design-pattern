package state;

public class On implements State{
    @Override
    public void switchOn(Switcher switcher) {
        System.out.println("WARN!!!通电状态无需再开");
    }

    @Override
    public void switchOff(Switcher switcher) {
        switcher.setState(new Off());
        System.out.println("OK...灯灭");
    }
}
