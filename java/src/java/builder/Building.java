package builder;

import java.util.ArrayList;
import java.util.List;

public class Building {

    private final List<String> buildingComponents = new ArrayList<>();

    public void setBasement(String basement) {
        this.buildingComponents.add(basement);
    }

    public void setWall(String wall) {
        this.buildingComponents.add(wall);
    }

    public void setRoof(String roof) {
        this.buildingComponents.add(roof);
    }

    @Override
    public String toString() {
        StringBuilder str = new StringBuilder();
        for (int i = buildingComponents.size() - 1; i >= 0; i--) {
            str.append(buildingComponents.get(i));
        }
//        for (String item: buildingComponents) {
//            str.append(item);
//        }
        return str.toString();
    }
}
