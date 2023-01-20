package ex02.java;

import java.util.Scanner;

class Ex02 {
    public static void main(String[] args) {
        System.out.print("What is the input string? ");
        Scanner scanner = new Scanner(System.in);
        String inputString = scanner.next();
        while (inputString.isEmpty()) {
            System.out.print("Please input anything! ");
            inputString = scanner.next();
        }
        System.out.println(inputString + " has " + inputString.length() + " characters.");
    }
}
