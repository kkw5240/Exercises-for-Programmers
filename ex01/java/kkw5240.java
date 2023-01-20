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
        return "Hello, " + name + ", " + getGreetingMessage(name);
    }

    private static String getGreetingMessage(String name) {
        GreetingType greetingType = GreetingType.valueOf(
                String.valueOf(
                        name.toUpperCase().charAt(0)
                )
        );

        return greetingType.getGreetingMessage();
    }

    enum GreetingType {
        A("good afternoon!"), B("nice to meet you!"), C("Come In!"), D("how was your day?"), E("good evening."), F("have fun~"), G("good to see you!"),
        H("have a good day!"), I("how are you these days?"), J("jolly your day."), K("what has kept you so busy?"), L("long time no see!"), M("good morning!"), N("good night!"),
        O("how is your business going?"), P("see you again."), Q("how are you doing?"), R("how are you?"), S("what a small world!"), T("Hello"), U("fancy meeting you here!"),
        V("wakanda forever!"), W("what's up?"), X("xxxx"), Y("Hello"), Z("i never expected to see you here!");

        private final String greetingMessage;

        GreetingType(String greetingMessage) {
            this.greetingMessage = greetingMessage;
        }

        public String getGreetingMessage() {
            return this.greetingMessage;
        }
    }
}