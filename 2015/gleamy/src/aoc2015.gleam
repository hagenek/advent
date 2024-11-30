// main.gleam
import aoc_types.{type DaySolution, DaySolution}
import days/day01
import days/day02
import days/day03
import gleam/int
import gleam/io
import gleam/list
import gleam/string as str
import simplifile
import utils.{format_answer}

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
  ]

  utils.solve_with_real_data(3, day03.part1)

  list.each(solutions, fn(day_solution) {
    let #(day, solution) = day_solution
    run_day(day, solution)
  })
}
