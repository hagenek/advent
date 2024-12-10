import day01
import day02
import gleam/io
import gleeunit
import gleeunit/should
import types.{Answer}

pub fn main() {
  gleeunit.main()
}

pub fn day1_part1_single_step_test() {
  let assert Answer(ans) =
    day01.part1(
      "3   4
4   3
2   5
1   3
3   9
3   3",
    )

  io.debug(ans)
  Answer(ans) |> should.equal(Answer(11))
}

// 33933 18696 diff = 15237
// 35446 34443 diff = 1003
// 46314 66062 diff = 19748
// 83974 83974 diff = 0
// Total diff = 35988

pub fn day01_part1_multi_test() {
  let assert Answer(ans) =
    day01.part1(
      "35446   18696
46314   66062
33933   83974
83974   34443",
    )

  io.debug(ans)
  Answer(ans) |> should.equal(Answer(35_988))
}

pub fn day01_part2_single_test() {
  let assert Answer(ans) =
    day01.part2(
      "3   4
4   3
2   5
1   3
3   9
3   3",
    )

  io.debug(ans)
  Answer(ans) |> should.equal(Answer(31))
}

// 7 6 4 2 1
// 1 2 7 8 9
// 9 7 6 2 1
// 1 3 2 4 5
// 8 6 4 4 1
// 1 3 6 7 9

// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.
pub fn day02_part1_single_test() {
  day02.line_ok([7, 6, 4, 2, 1]) |> should.be_true()
}

pub fn day02_part1_multi_test() {
  let assert Answer(ans) =
    day02.part1(
      "7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9",
    )

  io.debug(ans)
  Answer(ans) |> should.equal(Answer(2))
}

pub fn day02_part2_multi_test() {
  let assert Answer(ans) =
    day02.part2(
      "7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9",
    )

  io.debug(ans)
  Answer(ans) |> should.equal(Answer(4))
}
