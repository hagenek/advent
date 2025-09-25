# Advent of Code 2015 - Java 21

Enkel struktur for å løse Advent of Code 2015 oppgaver med Java 21 og Maven.

## Struktur

- `src/main/java/no/hagenek/aoc2015/` - Løsningsklasser
- `src/test/java/no/hagenek/aoc2015/` - Enhetstester
- `src/main/resources/` - Input-filer

## Bruk

1. **Kjør testene**: `mvn test`
2. **Kjør en løsning**: `mvn compile exec:java -Dexec.mainClass="no.hagenek.aoc2015.Day01"`

## Lage en ny dag

1. Kopier `DayTemplate.java` til `DayXX.java`
2. Kopier `DayTemplateTest.java` til `DayXXTest.java`
3. Opprett `src/main/resources/dayXX.txt` med din input
4. Implementer metodene basert på testene

## Eksempel: Dag 1

- `Day01.java` - Komplett løsning for dag 1
- `Day01Test.java` - Tester med eksempler fra problemteksten
- Kjør: `mvn test -Dtest=Day01Test` 