package ex06.java;

import java.time.LocalDate;
import java.util.Scanner;

class Ex06 {

    private final static int MIN_LEFT_YEARS = 0;

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        System.out.print("What is your current age? ");
        int age = scanner.nextInt();

        System.out.print("At what age would you like to retire? ");
        int retirementAge = scanner.nextInt();

        int thisYear = getThisYear();

        int leftYears = getLeftYears(age, retirementAge);
        if (isInvalidLeftYears(leftYears)) {
            System.out.println("You were already retired at " + (thisYear + leftYears) + ".");

            scanner.close();
            return;
        }

        System.out.println("You have " + leftYears + " years left until you can retire.");
        System.out.println("It's " + thisYear + ", so you can retire in " + (thisYear + leftYears) + ".");

        scanner.close();
    }

    private static int getLeftYears(int age, int retirementAge) {
        return retirementAge - age;
    }

    private static int getThisYear() {
        return LocalDate.now().getYear();
    }

    private static boolean isInvalidLeftYears(int leftYears) {
        return leftYears < MIN_LEFT_YEARS;
    }
}