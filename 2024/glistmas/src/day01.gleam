import gleam/dict
import gleam/int
import gleam/list
import gleam/string
import types.{type Solution, Answer}

fn extract_lines_from_data(data: String) -> List(String) {
  data
  |> string.trim
  |> string.split("\n")
  |> list.map(string.trim)
  |> list.filter(fn(line) { line != "" })
}

fn parse_to_columns(data: String) -> #(List(Int), List(Int)) {
  let lines = extract_lines_from_data(data)

  let parse_number = fn(str) {
    case int.parse(string.trim(str)) {
      Ok(i) -> i
      Error(_) -> panic as "Could not parse integer"
    }
  }

  let #(col_one, col_two) =
    lines
    |> list.map(fn(line) {
      let parts =
        line
        |> string.split(" ")
        |> list.filter(fn(part) { part != "" })

      case parts {
        [first, second] -> #(parse_number(first), parse_number(second))
        _ -> {
          panic as "Invalid line format"
        }
      }
    })
    |> list.unzip

  #(list.sort(col_one, int.compare), list.sort(col_two, int.compare))
}

pub fn count_frequencies(numbers: List(Int)) -> dict.Dict(Int, Int) {
  list.fold(numbers, dict.new(), fn(acc, num) {
    case dict.get(acc, num) {
      Ok(count) -> dict.insert(acc, num, count + 1)
      Error(_) -> dict.insert(acc, num, 1)
    }
  })
}

pub fn part1(data: String) -> Solution {
  let #(col_one, col_two) = parse_to_columns(data)

  list.zip(col_one, col_two)
  |> list.fold(0, fn(acc, pair) {
    let #(a, b) = pair
    let diff = int.absolute_value(b - a)
    acc + diff
  })
  |> Answer
}

pub fn part2(data: String) -> Solution {
  let #(left, right) = parse_to_columns(data)
  let right_frequencies = count_frequencies(right)

  left
  |> list.fold(0, fn(acc, num) {
    case dict.get(right_frequencies, num) {
      Ok(count) -> acc + num * count
      Error(_) -> acc
    }
  })
  |> Answer
}
