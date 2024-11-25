// test/aoc_test.gleam
import aoc_types.{Answer}
import days/day03
import gleeunit
import gleeunit/should

pub fn main() {
  gleeunit.main()
}

// Day 1 Tests
pub fn day01_part1_simple_test() {
  day03.part1(">")
  |> should.equal(Answer(2))
}
