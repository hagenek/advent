package no.hagenek.aoc2015;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

public class Day01 {
    
    public int solvePart1(String input) {
        return 0;
    }
    
    public int solvePart2(String input) {
        // TODO: Implementer del 2
        return 0;
    }
    
    public static String readInput() throws IOException {
        return Files.readString(Path.of("src/main/resources/day01.txt")).trim();
    }
    
    public static void main(String[] args) throws IOException {
        Day01 day = new Day01();
        String input = readInput();
        
        System.out.println("Part 1: " + day.solvePart1(input));
        System.out.println("Part 2: " + day.solvePart2(input));
    }
} 