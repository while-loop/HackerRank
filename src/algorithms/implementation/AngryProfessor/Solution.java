package algorithms.implementation.AngryProfessor;

import java.util.Scanner;

class Solution {

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);

        int T = in.nextInt();

        for (int trial = 0; trial < T; trial++) { // test cases
            for (int i = 0; i < 2; i++) { // input for each test case
                int N = in.nextInt();
                int K = in.nextInt();
                int[] times = new int[N];
                for (int j = 0; j < N; j++) {
                    times[j] = in.nextInt();
                }
                System.out.println(cancelClass(N, K, times));
            }
        }
    }

    private static String cancelClass(int n, int k, int[] times) {
        int count = 0;
        for (int i = 0; i < n; i++) {
            if (times[i] <= 0)
                count++;
            if (count >= k) {
                return "NO";
            }
        }

        return "YES";
    }
}