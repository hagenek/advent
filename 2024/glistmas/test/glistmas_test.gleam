import day01
import gleeunit
import gleeunit/should
import types.{Answer}

pub fn main() {
  gleeunit.main()
}

pub fn day03_part1_single_step_test() {
  day01.part1("hello")
  |> should.equal(Answer(5))
}
