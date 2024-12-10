import gleam/int
import gleam/io
import gleam/list
import gleam/string
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

pub fn delimited_list(
  str: String,
  split_on delimiter: String,
  using f: fn(String) -> a,
) -> List(a) {
  str |> string.split(delimiter) |> list.map(f)
}

pub fn ints(str: String, split_on delimiter: String) -> List(Int) {
  delimited_list(str, delimiter, int)
}
