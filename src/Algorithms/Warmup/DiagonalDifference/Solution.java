package algorithms.warmup.DiagonalDifference;

import java.util.Scanner;

public class Solution {

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);

        int size = in.nextInt();

        int[][] mat = new int[size][size];

        for (int i = 0; i < size; i++) {
            for (int j = 0; j < size; j++) {
                mat[i][j] = in.nextInt();
            }
        }

        int sum1 = 0, sum2 = sum1;
        for (int i = 0; i < size; i++) {
            sum1 += mat[i][i];
            sum2 += mat[i][size-1-i];
        }

        System.out.println(Math.abs(sum2-sum1));
    }
}