// test/aoc_test.gleam
import aoc_types.{Answer}
import days/day03
import gleeunit
import gleeunit/should

pub fn main() {
  gleeunit.main()
}

pub fn day03_part1_single_step_test() {
  day03.part1(">")
  |> should.equal(Answer(2))
}

pub fn day03_part1_multiple_step_test() {
  day03.part1("^>v<")
  |> should.equal(Answer(4))
}

pub fn day03_part1_zigzag_test() {
  day03.part1("^v^v^v^v^v")
  |> should.equal(Answer(2))
}

pub fn day03_part2_single_step_test() {
  day03.part2(">v")
  |> should.equal(Answer(3))
}

pub fn day03_part2_multiple_step_test() {
  day03.part2("^>v<")
  |> should.equal(Answer(3))
}

pub fn day03_part2_zigzag_test() {
  day03.part2("^v^v^v^v^v")
  |> should.equal(Answer(11))
}
