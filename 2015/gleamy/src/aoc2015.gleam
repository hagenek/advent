import days/day01
import gleam/int
import gleam/io
import simplifile

pub fn main() {
  let path = "./txt/day01.txt"
  let data = simplifile.read(path)

  let day1_part1 = case data {
    Ok(data) -> day01.part1(data)
    Error(_) -> panic as { "Couldnt read file at path " <> path }
  }
  let day1_part2 = case data {
    Ok(data) -> day01.part2(data)
    Error(_) -> panic as { "Couldnt read file at path " <> path }
  }

  io.println("Day1 part1 " <> int.to_string(day1_part1))
  io.println("Day1 part2 " <> int.to_string(day1_part2))
}
