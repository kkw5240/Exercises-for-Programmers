package ex07.java;

import java.util.Scanner;

class Ex07 {
    private final static double SQUARE_FEET_TO_SQUARE_METER = 0.09290304;

    public static void main(String[] args) {
        double length = promptLength();
        double width = promptWidth();

        double areaSquareFeet = getArea(length, width);
        double areaSquareMeter = convertFeetToMeter(areaSquareFeet);

        System.out.printf("""
                You entered dimensions of %.1f feet by %.1f feet
                The area is
                %.3f square feet
                %.3f square meters.
                %n""", length, width, areaSquareFeet, areaSquareMeter);
    }

    private static double promptLength() {
        Scanner scanner = new Scanner(System.in);

        System.out.print("What is the length of the room in feet? ");
        double length = scanner.nextDouble();

        scanner.close();

        return length;
    }

    private static double promptWidth() {
        Scanner scanner = new Scanner(System.in);

        System.out.print("What is the width of the room in feet? ");
        double width = scanner.nextDouble();

        scanner.close();

        return width;
    }

    private static double getArea(double length, double width) {
        return length * width;
    }

    private static double convertFeetToMeter(double areaSquareFeet) {
        return areaSquareFeet * SQUARE_FEET_TO_SQUARE_METER;
    }
}