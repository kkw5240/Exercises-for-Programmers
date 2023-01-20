package ex02.java;

import java.util.Scanner;

class Ex02 {
    public static void main(String[] args) {
        printPrompt();

        printLengthOf(
                getInputString()
        );
    }

    private static void printPrompt() {
        System.out.print("What is the input string? ");
    }

    private static void printLengthOf(String inputString) {
        System.out.println(inputString + " has " + inputString.length() + " characters.");
    }

    private static String getInputString() {
        Scanner scanner = new Scanner(System.in);

        String inputString = scanner.nextLine();
        while (inputString.isEmpty()) {
            System.out.print("Please input anything! ");
            inputString = scanner.nextLine();
        }

        scanner.close();

        return inputString;
    }
}
