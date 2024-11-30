import gleam/string
import types.{type Solution, Answer}

pub fn part1(data: String) -> Solution {
  string.length(data)
  |> Answer
}

pub fn part2(_data: String) -> Solution {
  Answer(0)
}
