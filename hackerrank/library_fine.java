package algorithms.warmup.LibraryFine;

import java.util.Scanner;

public class Solution {

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);

        String actual = in.nextLine();
        String expected = in.nextLine();

        System.out.println(calculateFine(actual.split(" "), expected.split(" ")) + "");
    }

    private static int calculateFine(String[] actual, String[] expected) {

        int fine = 0;
        int years = actual[2].compareTo(expected[2]);
        if (years > 0) {
            fine = 10000;
        } else if (years == 0) {
            int days = Integer.parseInt(actual[0]) - Integer.parseInt(expected[0]);
            int months = Integer.parseInt(actual[1]) - Integer.parseInt(expected[1]);

            if (months > 0) {
                days = 0;
            } else if (months < 0) {
                days = 0;
                months = 0;
            } else if (days <= 0) {
                months = 0;
                days = 0;
            }

            fine = (500 * months) + (15 * days);
        }
        return fine;
    }

}