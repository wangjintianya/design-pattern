package iterator;

import java.nio.charset.StandardCharsets;

public class DrivingRecorder {
    private int index = -1;
    private final String[] records = new String[10];

    public void append(String record) {
        if (index == 9) {
            index = 0;
        } else {
            index++;
        }
        records[index] = record;
    }

    public void display() {
        for (int i = 0; i < 10; i++) {
            System.out.println(records[i]);
        }
    }

    public void displayInOrder() {
        for (int i = index, loopCount = 0; loopCount < 10; i = i == 0 ? 9 : i - 1, loopCount++) {
            System.out.println(records[i]);
        }
    }

    private class Itr implements Iterator<String> {
        private int cursor = index;
        private int loopCount = 0;

        @Override
        public String next() {
            int i = cursor; // 记录即将返回的游标位置
            if (cursor == 0) {
                cursor = 9;
            } else {
                cursor--;
            }
            loopCount++;
            return records[i];
        }

        @Override
        public boolean hasNext() {
            return loopCount < 10;
        }
    }

    public Iterator<String> iterator() {
        return new Itr();
    }
}
