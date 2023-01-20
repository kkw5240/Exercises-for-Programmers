package ex03.java;

import java.util.Scanner;

class Ex03 {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        System.out.print("What is the quote? ");
        String quote = scanner.nextLine();

        System.out.print("Who said it? ");
        String person = scanner.nextLine();

        System.out.println(person + " says, \"" + quote + ".\"");
    }
}
