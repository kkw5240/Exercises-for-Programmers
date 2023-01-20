package ex03.java;

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

class Ex03 {
    private static Scanner scanner;
    public static void main(String[] args) {
        scanner = new Scanner(System.in);

        Quotes quotes = new Quotes();

        while (isEmptyInput()) {
            Quote quote = makeQuote();
            quote.printQuote();

            quotes.add(quote);
        }
        quotes.printQuotes();

        scanner.close();
    }

    private static Quote makeQuote() {
        String message = getMessage();
        String person = getPerson();

        return new Quote(person, message);
    }

    private static boolean isEmptyInput() {
        return scanner.nextLine().isEmpty();
    }

    private static String getMessage() {
        System.out.print("What is the quote? ");
        return scanner.nextLine();
    }

    private static String getPerson() {
        System.out.print("Who said it? ");
        return scanner.nextLine();
    }

    private static class Quote {
        private final String person;
        private final String message;

        public Quote(String person, String message) {
            this.person = person;
            this.message = message;
        }

        public void printQuote() {
            System.out.println(this.person + " says, \"" + this.message + ".\"");
        }
    }

    private static class Quotes {
        private final List<Quote> quotes;

        public Quotes() {
            this.quotes = new ArrayList<>();
        }

        public void add(Quote quote) {
            this.quotes.add(quote);
        }

        public void printQuotes() {
            this.quotes.forEach(Quote::printQuote);
            System.out.println();
        }
    }
}
