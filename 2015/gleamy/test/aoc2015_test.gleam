import days/day01
import gleeunit
import gleeunit/should

pub fn main() {
  gleeunit.main()
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
