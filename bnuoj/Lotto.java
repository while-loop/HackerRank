import java.io.*;
import java.util.Arrays;
import java.util.StringTokenizer;

// PID: 1700
// https://www.bnuoj.com/v3/problem_show.php?pid=1108
class Lotto {
    private final static FastIO io = new FastIO();

    public static void main(String[] args) {

        int n = io.nextInt();
        while (n != 0) {
            int[] list = new int[n];
            for (int i = 0; i < n; i++) {
                list[i] = io.nextInt();
            }
            System.out.println(Arrays.toString(list));

            combination(list, new int[6], 0, 0);
            io.println();
        }
    }

    private static void combination(int[] ogList, int[] curList, int ogIdx, int curIdx) {
        if (curIdx == curList.length) {
            io.println(curList);
            return;
        }

        for (int i = ogIdx; i < ogList.length; i++) {
            curList[curIdx] = ogList[i];
            combination(ogList, curList, ogIdx + 1, curIdx + 1);
        }
    }

    static class FastIO {
        private final BufferedReader br;
        private StringTokenizer st;
        private final PrintWriter writer;


        public FastIO() {
            this(System.in, System.out);
        }

        public FastIO(InputStream in, OutputStream out) {
            this.br = new BufferedReader(new InputStreamReader(in));
            this.writer = new PrintWriter(new BufferedWriter(new OutputStreamWriter(out)));
        }

        String next() {
            while (st == null || !st.hasMoreElements()) {
                try {
                    st = new StringTokenizer(br.readLine());
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
            return st.nextToken();
        }

        int nextInt() {
            return Integer.parseInt(next());
        }

        long nextLong() {
            return Long.parseLong(next());
        }

        double nextDouble() {
            return Double.parseDouble(next());
        }

        String nextLine() {
            String str = "";
            try {
                str = br.readLine();
            } catch (IOException e) {
                e.printStackTrace();
            }
            return str;
        }

        public void print(int[] objects) {
            if (objects.length > 0) writer.print(objects[0]);
            for (int i = 1; i < objects.length; i++) {
                writer.print(" " + objects[i]);
            }
        }

        public void println(int[] objects) {
            System.out.println("sdf" + Arrays.toString(objects));
            print(objects);
            writer.println();
        }

        public void println() {
            writer.println();
        }

        public void close() {
            try {
                writer.close();
                br.close();
            } catch (IOException e) {
                e.printStackTrace();
            }
        }

        public void flush() {
            writer.flush();
        }
    }
}
