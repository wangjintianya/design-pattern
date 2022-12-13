package iterator;

import java.util.ArrayList;
import java.util.List;

public class Client {
    public static void main(String[] args) {
        DrivingRecorder drivingRecorder = new DrivingRecorder();
        for (int i = 0; i < 12; i++) {
            drivingRecorder.append("第" + i + "视频");
        }

        List<String> accidents = new ArrayList<>();
        Iterator<String> iterator = drivingRecorder.iterator();
        while (iterator.hasNext()) {
            String record = iterator.next();
            if (record.equals("第3视频") || record.equals("第6视频")) {
                accidents.add(record);
            }
        }
        accidents.forEach(System.out::println);
    }
}
