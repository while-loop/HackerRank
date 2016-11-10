package algorithms.warmup.ExtraLongFactorials;

import java.math.BigInteger;
import java.util.Scanner;

public class Solution {

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);

        System.out.println(fact(in.nextInt()));
    }

    private static BigInteger fact(int i) {
        BigInteger factorial = BigInteger.ONE;
        for (int j = 1; j <= i; j++) {
           factorial = factorial.multiply(BigInteger.valueOf(j));
        }
        return factorial;
    }

    private static BigInteger recFact(int i) {
        if (i == 1) {
            return BigInteger.ONE;
        }
        return BigInteger.valueOf(i).multiply(recFact(i - 1));
    }
}