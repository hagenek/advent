import days/day01
import days/day02
import gleeunit
import gleeunit/should

pub fn main() {
  gleeunit.main()
}

// For example:

// A present with dimensions 2x3x4 requires 2+2+3+3 = 10
// feet of ribbon to wrap the present plus 2*3*4 = 24 feet of ribbon for the bow, for a total of 34 feet.
// A present with dimensions 1x1x10 requires 1+1+1+1 = 4
// feet of ribbon to wrap the present plus 1*1*10 = 10 feet of ribbon for the bow, for a total of 14 feet.

pub fn day02_part2_test() {
  day02.parse_line_ribbon("2x3x4")
  |> should.equal(34)

  day02.parse_line_ribbon("1x1x10")
  |> should.equal(14)
}

pub fn day02_init_test() {
  day02.parse_line_wrapping("2x3x4")
  |> should.equal(58)

  day02.part1("2x3x4\n1x1x10")
  |> should.equal(101)
}

// gleeunit test functions end in `_test`
pub fn part1_test() {
  day01.parse_line("(((")
  |> should.equal(3)
}

// gleeunit test functions end in `_test`
pub fn part1_both_test() {
  day01.parse_line("((()))(")
  |> should.equal(1)
}

pub fn part1_both_wrong_test() {
  day01.parse_line("((()))(((((")
  |> should.equal(5)
}
