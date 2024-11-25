// day02.gleam
import aoc_types.{type Solution, Answer}
import gleam/int
import gleam/list
import gleam/string

pub type Box {
  Box(length: Int, width: Int, height: Int)
}

pub fn find_smallest(a: Int, b: Int, c: Int) -> Int {
  let smallest_one = int.min(a, b)
  int.min(smallest_one, c)
}

pub fn calc_wrapping_paper(box: Box) -> Int {
  case box {
    Box(l, w, h) -> {
      let top_bot = l * w
      let front_back = w * h
      let side = h * l
      top_bot
      * 2
      + front_back
      * 2
      + side
      * 2
      + find_smallest(top_bot, front_back, side)
    }
  }
}

pub fn parse_line_wrapping(input: String) -> Int {
  let split_string = input |> string.split("x")
  case split_string {
    [l, w, h] -> {
      let assert Ok(length) = int.parse(l)
      let assert Ok(width) = int.parse(w)
      let assert Ok(height) = int.parse(h)
      calc_wrapping_paper(Box(length, width, height))
    }
    _ -> 0
  }
}

pub fn calc_ribbon_paper(box: Box) -> Int {
  case box {
    Box(l, w, h) -> {
      let shortest_distance = int.min(2 * w + 2 * h, 2 * w + 2 * l)
      let shortest_perimeter =
        find_smallest(2 * w + 2 * h, 2 * w + 2 * l, 2 * l + 2 * h)
      let total = int.min(shortest_distance, shortest_perimeter) + { l * w * h }
      total
    }
  }
}

pub fn parse_line_ribbon(input: String) -> Int {
  let split_string = input |> string.split("x")
  case split_string {
    [l, w, h] -> {
      let assert Ok(length) = int.parse(l)
      let assert Ok(width) = int.parse(w)
      let assert Ok(height) = int.parse(h)
      calc_ribbon_paper(Box(length, width, height))
    }
    _ -> 0
  }
}

pub fn part1(input: String) -> Solution {
  let result =
    input
    |> string.split("\n")
    |> list.map(parse_line_wrapping)
    |> list.fold(0, fn(a: Int, b: Int) { a + b })
  Answer(result)
}

pub fn part2(input: String) -> Solution {
  let result =
    input
    |> string.split("\n")
    |> list.map(parse_line_ribbon)
    |> list.fold(0, fn(a: Int, b: Int) { a + b })
  Answer(result)
}
