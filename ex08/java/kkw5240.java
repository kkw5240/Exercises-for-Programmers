package ex08.java;

import java.util.Scanner;

class Ex08 {
    public static void main(String[] args) {
        int people = 1, pizzas = 0, pieces = 0;

        Scanner scanner = new Scanner(System.in);

        try {
            System.out.printf("How many people? ");
            people = scanner.nextInt();

            System.out.printf("How many pizzas do you have? ");
            pizzas = scanner.nextInt();

            System.out.printf("How many pieces are in a pizza? ");
            pieces = scanner.nextInt();

        } catch (Exception e) {
            System.out.println(e.getMessage());
        } finally {
            scanner.close();
        }

        System.out.printf("\n%d people with %d pizzas.\n", people, pizzas);

        int sharedPieces = pizzas * pieces / people;
        System.out.printf("Each person gets %d pieces.\n", sharedPieces);

        int leftoverPieces = pizzas * pieces % people;
        System.out.printf("There are %d leftover pieces.\n", leftoverPieces);
    }
}