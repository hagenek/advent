// test/aoc_test.gleam
import aoc_types.{Answer}
import days/day01
import days/day02
import gleeunit
import gleeunit/should

pub fn main() {
  gleeunit.main()
}

// Day 1 Tests
pub fn day01_part1_simple_test() {
  day01.part1("(((")
  |> should.equal(Answer(3))
}

pub fn day01_part1_mixed_test() {
  day01.part1("((()))(")
  |> should.equal(Answer(1))
}

pub fn day01_part1_complex_test() {
  day01.part1("((()))(((((")
  |> should.equal(Answer(5))
}

pub fn day01_internal_parse_test() {
  day01.parse_line("(((")
  |> should.equal(3)
}

// Day 2 Tests
pub fn day02_part1_single_box_test() {
  day02.part1("2x3x4")
  |> should.equal(Answer(58))
}

pub fn day02_part1_multiple_boxes_test() {
  day02.part1("2x3x4\n1x1x10")
  |> should.equal(Answer(101))
}

pub fn day02_part2_single_box_test() {
  day02.part2("2x3x4")
  |> should.equal(Answer(34))
}

pub fn day02_part2_multiple_boxes_test() {
  day02.part2("2x3x4\n1x1x10")
  |> should.equal(Answer(48))
}

// Internal function tests
pub fn day02_internal_wrapping_test() {
  day02.parse_line_wrapping("2x3x4")
  |> should.equal(58)

  day02.parse_line_wrapping("1x1x10")
  |> should.equal(43)
}

pub fn day02_internal_ribbon_test() {
  day02.parse_line_ribbon("2x3x4")
  |> should.equal(34)

  day02.parse_line_ribbon("1x1x10")
  |> should.equal(14)
}

// Helper function tests
pub fn day02_box_wrapping_test() {
  day02.calc_wrapping_paper(day02.Box(2, 3, 4))
  |> should.equal(58)

  day02.calc_wrapping_paper(day02.Box(1, 1, 10))
  |> should.equal(43)
}

pub fn day02_box_ribbon_test() {
  day02.calc_ribbon_paper(day02.Box(2, 3, 4))
  |> should.equal(34)

  day02.calc_ribbon_paper(day02.Box(1, 1, 10))
  |> should.equal(14)
}
