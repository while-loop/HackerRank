package algorithms.warmup.AVeryBigSum;

import java.util.Scanner;

class Solution {

    public static void main(String[] args) {
        /* Enter your code here. Read input from STDIN. Print output to STDOUT. Your class should be named Solution. */
        Scanner in = new Scanner(System.in);

        int numbers = in.nextInt();

        int sum = 0;
        for (int i = 0; i < numbers; i++) {
            sum += in.nextInt();
        }

        System.out.println(sum);
    }
}