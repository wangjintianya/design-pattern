package builder;

public interface Builder {
    public void buildBasement();
    public void buildWall();
    public void buildRoof();

    public Building getBuilding();
}
