package ex01.java;

import java.util.Scanner;

/**
 * 작성자 : 김규완
 * 실행방법 : java kkw5240.class
 * 주의사항 : 기본풀이까지만 했습니다.
 */
class Ex01 {
    public static void main(String[] args) {
        printPrompt();

        Scanner scanner = new Scanner(System.in);
        if (isScanned(scanner)) {
            printHello(scanner.next());
        }
    }

    private static void printPrompt() {
        System.out.print("What is your name? ");
    }

    private static boolean isScanned(Scanner scanner) {
        return scanner.hasNext();
    }

    private static void printHello(String name) {
        System.out.println(getMessage(name));
    }

    private static String getMessage(String name) {
        return "Hello, " + name + ", nice to meet you!";
    }
}