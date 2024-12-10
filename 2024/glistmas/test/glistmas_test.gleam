import day01
import day03
import day04
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

pub fn day04_part_1_test() {
  "MMMSXXMASM"
  |> day04.part1
  |> should.equal(types.Answer(1))
}

pub fn day04_part_1_multiple_lines_test() {
  "MMMSXXMASM\nMMMSXXMASM"
  |> day04.part1
  |> should.equal(types.Answer(2))
}

pub fn day04_part_1_multiple_lines_backwards_test() {
  "MMMSXXSAMM\nMMSAMXXSAMM"
  |> day04.part1
  |> should.equal(types.Answer(1))
}

pub fn day04_part_1_multiple_lines_vertical_test() {
  "XMMXXXSAMM
MMSAXXXSAMM
AMSAXXXSAMM
SMSAXXXSAMM
"
  |> day04.part1
  |> should.equal(types.Answer(1))
}

pub fn day04_part_1_multiple_lines_diagonal_test() {
  "X*******
*M*********
**A********
***S*******
"
  |> day04.part1
  |> should.equal(types.Answer(1))
}

pub fn day04_part_1_multiple_lines_diagonal_2_test() {
  "X******X***
*M******M**
**A******A*
***S******S
"
  |> day04.part1
  |> should.equal(types.Answer(2))
}

pub fn day04_part_1_multiple_lines_diagonal_mini_forward_backward_test() {
  "XScde
1MA45
fgAMh
lmnSX
"
  |> day04.part1
  |> should.equal(types.Answer(2))
}

pub fn day04_part_1_multiple_lines_diagonal_mini_forwadsfard_backward_test() {
  "
XbcdSe**
1MA45A**
fgAMh*M*
lmnSX**X
"
  |> day04.part1
  |> should.equal(types.Answer(2))
}

pub fn day04_part_1_multiple_lines_diagonal_6_test() {
  "X***S***
*M***A**
**A***M*
***S***X
"
  |> day04.part1
  |> should.equal(types.Answer(2))
}

pub fn day04_part_2_multiple_lines_diagonal_1_test() {
  let first =
    "M*S
*A*
M*S"

  let second =
    "M*M
*A*
S*S"

  let third =
    "S*S
*A*
M*M"

  third |> day04.part2 |> should.equal(types.Answer(1))
  second |> day04.part2 |> should.equal(types.Answer(1))
  first |> day04.part2 |> should.equal(types.Answer(1))
}

pub fn day04_part_2_multiple_lines_diagonal_2_test() {
  "
.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
.........."
  |> day04.part2
  |> should.equal(types.Answer(9))
}
