package state;

public class Switcher{

    private State state = new Off();

    public State getState() {
        return state;
    }

    public void setState(State state) {
        this.state = state;
    }

    public void switchOn() {
       state.switchOn(this);
    }


    public void switchOff() {
        state.switchOff(this);
    }
}
