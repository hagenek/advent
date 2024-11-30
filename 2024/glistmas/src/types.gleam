pub type Solution {
  Answer(Int)
  StringAnswer(String)
}

pub type DaySolution {
  DaySolution(part1: fn(String) -> Solution, part2: fn(String) -> Solution)
}
