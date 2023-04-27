package ex07.java;

import java.util.Scanner;

class Ex07 {
    private final static double SQUARE_FEET_TO_SQUARE_METER = 0.09290304;
    private double length;
    private double width;
    private double areaSquareFeet;
    private double areaSquareMeter;

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        System.out.print("What is the length of the room in feet? ");
        double length = scanner.nextDouble();

        System.out.print("What is the width of the room in feet? ");
        double width = scanner.nextDouble();

        double areaSquareFeet = length * width;
        double areaSquareMeter = areaSquareFeet * SQUARE_FEET_TO_SQUARE_METER;

        String message = String.format("""
                You entered dimensions of %.1f feet by %.1f feet
                The area is
                %.3f square feet
                %.3f square meters.
                """, length, width, areaSquareFeet, areaSquareMeter);

        System.out.println(message);
    }
}