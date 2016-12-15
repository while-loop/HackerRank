package algorithms.warmup.Staircase;

import java.util.Scanner;

public class Solution {

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);

        int N = in.nextInt();

        for (int i = 1; i <= N; i++) {
            for (int j = N; j > 0; j--) {
                if (j > i) {
                    System.out.print(" ");
                } else {
                    System.out.print("#");
                }
            }
            System.out.println();
        }

    }
}