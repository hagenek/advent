// day02.gleam
import aoc_types.{type Solution, Answer}
import gleam/bool
import gleam/io
import gleam/list
import gleam/string

pub type Direction {
  Up
  Down
  Left
  Right
}

pub type Coordinates {
  Coordinates(x: Int, y: Int)
}

pub type ParsingError {
  EmptyLineError
  UnknownCharError(char: String)
}

pub fn transform_to_direction(
  line: String,
) -> Result(List(Direction), ParsingError) {
  let split_string = line |> string.to_graphemes
  use <- bool.guard(list.is_empty(split_string), Error(EmptyLineError))

  split_string
  |> list.map(fn(char) {
    case char {
      ">" -> Up
      "^" -> Down
      "<" -> Left
      "v" -> Right
      _ -> panic as "Unknown char encountered"
    }
  })
  |> Ok
}

fn move_to_coordinates(
  direction: Direction,
  coordinates: Coordinates,
) -> Coordinates {
  case direction {
    Up -> Coordinates(x: coordinates.x, y: coordinates.y + 1)
    Down -> Coordinates(x: coordinates.x, y: coordinates.y - 1)
    Left -> Coordinates(x: coordinates.x - 1, y: coordinates.y)
    Right -> Coordinates(x: coordinates.x + 1, y: coordinates.y)
  }
}

// Visited nodes are stored in a list

pub fn part1(input: String) -> Solution {
  let direction_list = case transform_to_direction(input) {
    Ok(directions) -> directions
    _ -> panic as "unknown char encountered"
  }
  io.debug("Direction list: \n")
  io.debug(direction_list)
  Answer(0)
}

pub fn part2(input: String) -> Solution {
  todo
}
