// main.gleam
import aoc_types.{
  type DaySolution, type Solution, Answer, DaySolution, StringAnswer,
}
import days/day01
import days/day02
import gleam/int
import gleam/io
import gleam/list
import gleam/string as str
import simplifile

pub fn format_answer(answer: Solution) -> String {
  case answer {
    Answer(n) -> int.to_string(n)
    StringAnswer(s) -> s
  }
}

pub fn run_day(day: Int, solution: DaySolution) {
  let path = "./txt/day" <> str.pad_left(int.to_string(day), 2, "0") <> ".txt"

  case simplifile.read(path) {
    Ok(data) -> {
      let part1_result = solution.part1(data)
      let part2_result = solution.part2(data)

      io.println(
        "Day "
        <> int.to_string(day)
        <> " part 1: "
        <> format_answer(part1_result),
      )
      io.println(
        "Day "
        <> int.to_string(day)
        <> " part 2: "
        <> format_answer(part2_result),
      )
    }
    Error(_) -> panic as { "Couldn't read file at path " <> path }
  }
}

pub fn main() {
  let solutions = [
    #(1, DaySolution(part1: day01.part1, part2: day01.part2)),
    #(2, DaySolution(part1: day02.part1, part2: day02.part2)),
    // Add more days here in the same pattern
  ]

  list.each(solutions, fn(day_solution) {
    let #(day, solution) = day_solution
    run_day(day, solution)
  })
}
