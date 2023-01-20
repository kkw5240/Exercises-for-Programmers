package ex05.java;

import java.util.Scanner;

class Ex05 {
    static Scanner scanner;

    public static void main(String[] args) {
        scanner = new Scanner(System.in);

        System.out.print("What is the first number? ");
        NaturalNumber firstNumber = new NaturalNumber(scanner.nextLine());

        System.out.print("What is the second number? ");
        NaturalNumber secondNumber = new NaturalNumber(scanner.nextLine());

        System.out.println(
                firstNumber.getNumber() + " + " + secondNumber.getNumber() + " = " + NaturalNumber.add(firstNumber, secondNumber) + "\n"
                        + firstNumber.getNumber() + " - " + secondNumber.getNumber() + " = " + NaturalNumber.sub(firstNumber, secondNumber) + "\n"
                        + firstNumber.getNumber() + " * " + secondNumber.getNumber() + " = " + NaturalNumber.mul(firstNumber, secondNumber) + "\n"
                        + firstNumber.getNumber() + " / " + secondNumber.getNumber() + " = " + NaturalNumber.div(firstNumber, secondNumber)
        );

        scanner.close();
    }
}

class NaturalNumber {
    private final String number;

    public NaturalNumber(String number) {
        if (number.matches("^-\\d+|\\D+")) {
            throw new IllegalArgumentException();
        }

        this.number = number;
    }

    public int getNumber() {
        return Integer.parseInt(this.number);
    }

    public static int add(NaturalNumber firstNumber, NaturalNumber secondNumber) {
        return firstNumber.getNumber() + secondNumber.getNumber();
    }

    public static int sub(NaturalNumber firstNumber, NaturalNumber secondNumber) {
        return firstNumber.getNumber() - secondNumber.getNumber();
    }

    public static int mul(NaturalNumber firstNumber, NaturalNumber secondNumber) {
        return firstNumber.getNumber() * secondNumber.getNumber();
    }

    public static int div(NaturalNumber firstNumber, NaturalNumber secondNumber) {
        return firstNumber.getNumber() / secondNumber.getNumber();
    }
}