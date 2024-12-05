import day01
import day03
import gleam/io
import gleeunit
import gleeunit/should
import types

pub fn main() {
  gleeunit.main()
}

pub fn day1_part1_single_step_test() {
  let assert types.Answer(ans) =
    day01.part1(
      "3   4
4   3
2   5
1   3
3   9
3   3",
    )

  io.debug(ans)
  types.Answer(ans) |> should.equal(types.Answer(11))
}

// 33933 18696 diff = 15237
// 35446 34443 diff = 1003
// 46314 66062 diff = 19748
// 83974 83974 diff = 0
// Total diff = 35988

pub fn day01_part1_multi_test() {
  let assert types.Answer(ans) =
    day01.part1(
      "35446   18696
46314   66062
33933   83974
83974   34443",
    )

  io.debug(ans)
  types.Answer(ans) |> should.equal(types.Answer(35_988))
}

pub fn day01_part2_single_test() {
  let assert types.Answer(ans) =
    day01.part2(
      "3   4
4   3
2   5
1   3
3   9
3   3",
    )

  io.debug(ans)
  types.Answer(ans) |> should.equal(types.Answer(31))
}

// Test 1: Basic example from problem statement
pub fn day03_part1_test1() {
  ");>]mul(2,4)mul(3,3)"
  |> day03.part1
  |> should.equal(types.Answer(17))
}

pub fn day03_part1_test2() {
  "mul(5,5)mul(2,2)"
  |> day03.part1
  |> should.equal(types.Answer(29))
}

pub fn day03_part1_test3() {
  "mul(3,4)mul(2,5)mul(1,6)"
  |> day03.part1
  |> should.equal(types.Answer(32))
}

pub fn day03_part2_1_test() {
  "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)do()?mul(8,5))"
  |> day03.part2
  |> should.equal(types.Answer(48))
}

pub fn day03_part2_2_test() {
  "mul(2,3)don't()mul(4,5)mul(1,2)do()mul(3,4)"
  |> day03.part2
  |> should.equal(types.Answer(18))
}

pub fn day03_part2_3_test() {
  "don't()mul(5,5)do()mul(2,2)don't()mul(3,3)"
  |> day03.part2
  |> should.equal(types.Answer(4))
}
