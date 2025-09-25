package no.hagenek.aoc2015;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

public class DayTemplate {
    
    public String solvePart1(String input) {
        // TODO: Implementer del 1
        return "";
    }
    
    public String solvePart2(String input) {
        // TODO: Implementer del 2
        return "";
    }
    
    public static String readInput(String filename) throws IOException {
        return Files.readString(Path.of("src/main/resources/" + filename)).trim();
    }
    
    public static void main(String[] args) throws IOException {
        DayTemplate day = new DayTemplate();
        String input = readInput("dayXX.txt"); // Endre til riktig fil
        
        System.out.println("Part 1: " + day.solvePart1(input));
        System.out.println("Part 2: " + day.solvePart2(input));
    }
} 