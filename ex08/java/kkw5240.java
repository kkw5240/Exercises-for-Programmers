package ex08.java;

import java.util.Scanner;

class Ex08 {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        System.out.printf("How many people? ");
        int people = scanner.nextInt();

        System.out.printf("How many pizzas do you have? ");
        int pizzas = scanner.nextInt();

        System.out.printf("How many pieces are in a pizza? ");
        int pieces = scanner.nextInt();

        scanner.close();

        System.out.printf("\n%d people with %d pizzas.\n", people, pizzas);

        int sharedPieces = pizzas * pieces / people;
        System.out.printf("Each person gets %d pieces.\n", sharedPieces);

        int leftoverPieces = pizzas * pieces % people;
        System.out.printf("There are %d leftover pieces.\n", leftoverPieces);
    }
}