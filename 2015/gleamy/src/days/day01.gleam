// day01.gleam
import aoc_types.{type Solution, Answer}
import gleam/list.{Continue, Stop}
import gleam/string

fn paren_to_number(paren: String) -> Int {
  case paren {
    "(" -> 1
    ")" -> -1
    _ -> 0
  }
}

fn paren_to_number_indexed(paren: String, index: Int) -> #(Int, Int) {
  case paren {
    "(" -> #(1, index)
    ")" -> #(-1, index)
    _ -> #(0, index)
  }
}

fn fold_add_indexed(acc: Int, item: #(Int, Int)) -> list.ContinueOrStop(Int) {
  case acc + item.0 {
    -1 -> Stop(item.1)
    _ -> Continue(acc + item.0)
  }
}

fn fold_add(acc: Int, item: Int) -> Int {
  acc + item
}

pub fn parse_line(str: String) -> Int {
  str
  |> string.to_graphemes
  |> list.map(paren_to_number)
  |> list.fold(0, fold_add)
}

fn parse_line_index(str: String) -> Int {
  str
  |> string.to_graphemes
  |> list.index_map(paren_to_number_indexed)
  |> list.fold_until(1, fold_add_indexed)
}

pub fn part1(input: String) -> Solution {
  input
  |> parse_line
  |> Answer
}

pub fn part2(input: String) -> Solution {
  input
  |> parse_line_index
  |> Answer
}
