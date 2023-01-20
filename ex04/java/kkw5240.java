package ex04.java;

import java.util.Scanner;

class Ex04 {
    private static Scanner scanner;

    public static void main(String[] args) {
        scanner = new Scanner(System.in);

        MadLibs madLibs = MadLibs.builder()
                .noun(getNoun())
                .verb(getVerb())
                .adjective(getAdjective())
                .adverb(getAdverb())
                .build();

        System.out.println(madLibs);

        scanner.close();
    }

    private static String getAdverb() {
        System.out.print("Enter an adverb: ");
        return scanner.nextLine();
    }

    private static String getAdjective() {
        System.out.print("Enter an adjective: ");
        return scanner.nextLine();
    }

    private static String getVerb() {
        System.out.print("Enter a verb: ");
        return scanner.nextLine();
    }

    private static String getNoun() {
        System.out.print("Enter a noun: ");
        return scanner.nextLine();
    }
}

class MadLibs {
    private final String madLibs;

    private MadLibs(String noun, String verb, String adjective, String adverb) {
        this.madLibs = makeMadLibs(
                verb,
                adjective,
                noun,
                adverb
        );
    }

    private String makeMadLibs(String verb,
                               String adjective,
                               String noun,
                               String adverb) {
        return String.format(
                "Do you %s your %s %s %s? That's hilarious!",
                verb,
                adjective,
                noun,
                adverb
        );
    }

    public String toString() {
        return madLibs;
    }

    public static MadLibs.Builder builder() {
        return new MadLibs.Builder();
    }

    public static class Builder {
        private String noun;
        private String verb;
        private String adjective;
        private String adverb;

        public MadLibs.Builder noun(String noun) {
            this.noun = noun;
            return this;
        }

        public MadLibs.Builder verb(String verb) {
            this.verb = verb;
            return this;
        }

        public MadLibs.Builder adjective(String adjective) {
            this.adjective = adjective;
            return this;
        }

        public MadLibs.Builder adverb(String adverb) {
            this.adverb = adverb;
            return this;
        }

        public MadLibs build() {
            return new MadLibs(noun, verb, adjective, adverb);
        }
    }
}
