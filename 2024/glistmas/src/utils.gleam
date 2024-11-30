import gleam/int
import gleam/io
import gleam/string
import simplifile
import types.{type Solution, Answer, StringAnswer}

pub fn format_answer(answer: Solution) -> String {
  case answer {
    Answer(n) -> int.to_string(n)
    StringAnswer(s) -> s
  }
}

pub fn solve_with_real_data(
  day day: Int,
  with solution_function: fn(String) -> Solution,
) {
  let path =
    "./txt/day" <> string.pad_start(int.to_string(day), 2, "0") <> ".txt"

  case simplifile.read(path) {
    Ok(data) -> {
      let clean_data = string.trim(data)
      let part1_result = solution_function(clean_data)

      io.println(
        "Day "
        <> int.to_string(day)
        <> " part 1: "
        <> format_answer(part1_result),
      )
    }
    Error(_) -> panic as { "Couldn't read file at path " <> path }
  }
}
