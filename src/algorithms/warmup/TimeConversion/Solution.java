package algorithms.warmup.TimeConversion;

import java.util.Scanner;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Solution {

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);

        String hr12 = in.next();

        Pattern pattern = Pattern.compile("(\\d*):(\\d*:\\d*)([A|P]M)");
        Matcher matcher = pattern.matcher(hr12);
        if (!matcher.find()) {
            System.out.println("Unable to parse time");
            return;
        }

        int hour = Integer.parseInt(matcher.group(1));
        String mmss = matcher.group(2);
        String AMPM = matcher.group(3);


        if (AMPM.equals("PM") && hour != 12) {
                hour += 12;
        } else if (AMPM.equals("AM") && hour == 12) {
            hour = 0;
        }

        System.out.format("%02d:%s", hour, mmss);
    }
}