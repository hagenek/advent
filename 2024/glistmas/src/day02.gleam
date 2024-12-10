import gleam/list
import gleam/string
import types
import utils

@external(erlang, "list_variants", "remove_one")
pub fn remove_variants(input: List(Int)) -> List(List(Int))

pub fn parse(input: String) -> List(List(Int)) {
  let rows = {
    use line <- list.map(string.split(input, "\n"))
    let numbers_as_strings = string.split(line, " ")
    list.map(numbers_as_strings, utils.int)
  }
  rows
}

fn all_increasing(line: List(Int)) {
  line
  |> list.window_by_2
  |> list.map(fn(int_pair: #(Int, Int)) {
    let #(a, b) = int_pair
    b > a && b - a <= 3
  })
}

pub fn line_ok(line: List(Int)) -> Bool {
  let all_inc_list = all_increasing(line)
  let all_dec_list = all_increasing(line |> list.reverse)

  all_inc_list |> list.all(fn(x) { x }) || all_dec_list |> list.all(fn(x) { x })
}

pub fn part1(input: String) -> types.Solution {
  let lines = parse(input)
  let valid_lines = list.filter(lines, line_ok)
  list.length(valid_lines) |> types.Answer
}

pub fn part2(input: String) -> types.Solution {
  let lines = parse(input)
  let valid_count =
    list.filter(lines, fn(line) {
      let possible_lines = remove_variants(line)
      let valid_lines = list.filter(possible_lines, line_ok)
      list.length(valid_lines) > 0
    })
    |> list.length
  types.Answer(valid_count)
}
