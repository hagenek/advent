import gleam/list
import gleam/string

pub fn split_lines(with lines: String, on match_str: String) {
  lines
  |> string.trim
  |> string.split(match_str)
  |> list.map(string.trim)
  |> list.filter(fn(s) { !string.is_empty(s) })
}

pub fn split_on_newline(with lines: String) {
  split_lines(lines, "\n")
}
