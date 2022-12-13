package template;

public abstract class PM {
    public abstract void analyze();
    public abstract void design();
    public abstract void develop();
    public abstract boolean test();
    public abstract void release();

    public final void kickoff() {
        analyze();
        design();
        develop();
        while (test()) {
            develop();
        }
        release();
    }
}
