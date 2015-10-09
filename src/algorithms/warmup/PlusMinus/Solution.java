package algorithms.warmup.PlusMinus;

import java.util.Scanner;

public class Solution {

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);

        int N = in.nextInt();

        int pos = 0, neg = 0, zero = 0;
        for (int i = 0; i < N; i++) {
            int input = in.nextInt();
            if (input > 0) {
                pos++;
            } else if (input < 0) {
                neg++;
            } else {
                zero++;
            }
        }

        String format = "%.3f%n";
        System.out.format(format, (double)pos/N);
        System.out.format(format, (double)neg/N);
        System.out.format(format, (double)zero/N);
    }
}