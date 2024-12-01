import day01
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
