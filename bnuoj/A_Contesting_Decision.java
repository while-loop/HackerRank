import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.StringTokenizer;

class Main {

    public static void main(String[] args) {
        FastReader in = new FastReader();
        int teams = in.nextInt();

        String bestTeam = "";
        int probsSovled = Integer.MIN_VALUE, points = Integer.MAX_VALUE;

        for (int i = 0; i < teams; i++) {
            int submissions = 0, pts = 0;
            String team = in.next();
            for (int j = 0; j < 4; j++) {
                int t1 = in.nextInt();
                int t2 = in.nextInt();

                if (t2 > 0) {
                    pts += t2 + ((t1 - 1) * 20);
                    submissions++;
                }
            }

            if (submissions > 0) {
                if (submissions > probsSovled) {
                    probsSovled = submissions;
                    points = pts;
                    bestTeam = team;
                } else if ((submissions == probsSovled) && (pts < points)) {
                    probsSovled = submissions;
                    points = pts;
                    bestTeam = team;
                }
            }
        }
        System.out.println(String.format("%s %d %d", bestTeam, probsSovled, points));
    }

    static class FastReader {
        BufferedReader br;
        StringTokenizer st;

        public FastReader() {
            br = new BufferedReader(new
                    InputStreamReader(System.in));
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
    }
}