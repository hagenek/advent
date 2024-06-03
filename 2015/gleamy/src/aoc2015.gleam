import days/day01
import days/day02
import gleam/int
import gleam/io
import gleam/result
import simplifile

pub fn main() {
  let path = "./txt/day01.txt"
  let data_01 = simplifile.read(path)

  let day1_part1 = case data_01 {
    Ok(data) -> day01.part1(data)
    Error(_) -> panic as { "Couldnt read file at path " <> path }
  }
  let day1_part2 = case data_01 {
    Ok(data) -> day01.part2(data)
    Error(_) -> panic as { "Couldnt read file at path " <> path }
  }

  io.println("Day1 part1 " <> int.to_string(day1_part1))
  io.println("Day1 part2 " <> int.to_string(day1_part2))

  let path_02 = "./txt/day02.txt"
  let data_02 = simplifile.read(path_02)

  let day02_part1 = case data_02 {
    Ok(data) -> day02.part1(data)
    Error(_) -> panic as { "Couldnt read file at path " <> path }
  }

  io.println("Day2 part 1: " <> int.to_string(day02_part1))

  let day02_part2 = case data_02 {
    Ok(data) -> day02.part2(data)
    Error(_) -> panic as { "Couldnt read file at path " <> path }
  }

  io.println("Day2 part 2: " <> int.to_string(day02_part2))
}
