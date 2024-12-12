import gleam/int
import gleam/io
import gleam/list
import gleam/string
import simplifile
import types.{type Solution, Answer, StringAnswer}

pub fn format_answer(answer: Solution) -> String {
  case answer {
    Answer(n) -> int.to_string(n)
    StringAnswer(s) -> s
  }
}

pub fn debug_return(x: a) -> a {
  io.debug(x)
  x
}

pub fn int(str: String) -> Int {
  let assert Ok(n) = int.parse(str)
  n
}

/// Takes a string, splits it at each delimiter, then transforms each section using the given function
pub fn delimited_list(
  str: String,
  split_on delimiter: String,
  using f: fn(String) -> a,
) -> List(a) {
  str |> string.split(delimiter) |> list.map(f)
}

/// Splits a string on the given delimiter and parses each section as an integer
pub fn ints(str: String, split_on delimiter: String) -> List(Int) {
  delimited_list(str, delimiter, int)
}

pub fn solve_with_real_data(
  day day: Int,
  with solution_function: fn(String) -> Solution,
  part part: String,
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
        <> " part "
        <> part
        <> " result: "
        <> format_answer(part1_result),
      )
    }
    Error(_) -> panic as { "Couldn't read file at path " <> path }
  }
}
