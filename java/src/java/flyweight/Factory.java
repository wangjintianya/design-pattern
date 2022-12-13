package flyweight;

import java.util.HashMap;
import java.util.Map;

public class Factory {
    private Map<String, Drawable> images;

    public Factory() {
        images = new HashMap<>();
    }

    public Drawable getDrawable(String image) {
        if (!images.containsKey(image)) {
            switch (image){
                case "河流":
                    images.put(image, new Water());
                    break;
                case "草坪":
                    images.put(image, new Grass());
                    break;
                case "石路":
                    images.put(image, new Stone());
                    break;
                case "房子":
                    images.put(image, new House());
                    break;
            }
        }

        return images.get(image);
    }
}
