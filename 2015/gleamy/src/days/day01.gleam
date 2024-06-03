import gleam/io
import gleam/list.{Continue, Stop}
import gleam/string

// --- Day 1: Not Quite Lisp ---
// Santa was hoping for a white Christmas, but his weather machine's "snow" function is powered by stars, and he's fresh out! To save Christmas, he needs you to collect fifty stars by December 25th.

// Collect stars by helping Santa solve puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

// Here's an easy puzzle to warm you up.

// Santa is trying to deliver presents in a large apartment building, but he can't find the right floor - the directions he got are a little confusing. He starts on the ground floor (floor 0) and then follows the instructions one character at a time.

// An opening parenthesis, (, means he should go up one floor, and a closing parenthesis, ), means he should go down one floor.

// The apartment building is very tall, and the basement is very deep; he will never find the top or bottom floors.

// For example:

// (()) and ()() both result in floor 0.
// ((( and (()(()( both result in floor 3.
// ))((((( also results in floor 3.
// ()) and ))( both result in floor -1 (the first basement level).
// ))) and )())()) both result in floor -3.
// To what floor do the instructions take Santa?

fn paren_to_number(paren: String) -> Int {
  case paren {
    "(" -> 1
    ")" -> -1
    _ -> 0
  }
}

fn paren_to_number_indexed(paren: String, index: Int) {
  case paren {
    "(" -> #(1, index)
    ")" -> #(-1, index)
    _ -> #(0, index)
  }
}

fn fold_add_indexed(acc: Int, item: #(Int, Int)) {
  io.debug(#(acc, item))
  case acc + item.0 {
    -1 -> Stop(item.1)
    _ -> Continue(acc + item.0)
  }
}

fn fold_add(acc: Int, item: Int) {
  acc + item
}

pub fn parse_line(str: String) -> Int {
  str
  |> string.to_graphemes
  |> list.map(paren_to_number)
  |> list.fold(0, fold_add)
}

pub fn parse_line_index(str: String) -> Int {
  str
  |> string.to_graphemes
  |> list.index_map(paren_to_number_indexed)
  |> list.fold_until(1, fold_add_indexed)
}

pub fn part1(input: String) -> Int {
  input
  |> parse_line
}

pub fn part2(input: String) -> Int {
  input |> parse_line_index
}
