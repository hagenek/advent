#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_LINE_LENGTH 1024

void part1(const char* filename) {
    FILE *file = fopen(filename, "r");
    if (!file) {
        fprintf(stderr, "Error: Could not open file %s\n", filename);
        return;
    }

    char line[MAX_LINE_LENGTH];
    int result = 0;

    while (fgets(line, sizeof(line), file)) {
        // Remove newline if present
        line[strcspn(line, "\n")] = 0;

        // TODO: Implement part 1 logic here
    }

    fclose(file);
    printf("Part 1: %d\n", result);
}

void part2(const char* filename) {
    FILE *file = fopen(filename, "r");
    if (!file) {
        fprintf(stderr, "Error: Could not open file %s\n", filename);
        return;
    }

    char line[MAX_LINE_LENGTH];
    int result = 0;

    while (fgets(line, sizeof(line), file)) {
        // Remove newline if present
        line[strcspn(line, "\n")] = 0;

        // TODO: Implement part 2 logic here
    }

    fclose(file);
    printf("Part 2: %d\n", result);
}

int main(int argc, char *argv[]) {
    const char *filename = (argc > 1) ? argv[1] : "input/day1.txt";

    printf("Advent of Code 2015 - Day 1\n");
    printf("Input file: %s\n\n", filename);

    part1(filename);
    part2(filename);

    return 0;
}