import day01
import day02
import gleam/int
import gleam/io
import gleam/string
import simplifile
import types
import utils.{format_answer}

pub fn solve_with_real_data(
  day day: Int,
  with solution_function: fn(String) -> types.Solution,
  part part: Int,
) {
  let path =
    "./txt/day" <> string.pad_start(int.to_string(day), 2, "0") <> ".txt"

  case simplifile.read(path) {
    Ok(data) -> {
      let clean_data = string.trim(data)
      let result = solution_function(clean_data)

      io.println(
        "Day "
        <> int.to_string(day)
        <> " part "
        <> int.to_string(part)
        <> ": "
        <> format_answer(result),
      )
    }
    Error(_) -> panic as { "Couldn't read file at path " <> path }
  }
}

pub fn main() {
  solve_with_real_data(1, day01.part1, 1)
  solve_with_real_data(1, day01.part2, 2)
  solve_with_real_data(2, day02.part1, 1)
  solve_with_real_data(2, day02.part2, 2)
}
