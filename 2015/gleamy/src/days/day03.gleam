// day02.gleam
import aoc_types.{type Solution, Answer}
import gleam/bool
import gleam/list
import gleam/result
import gleam/set.{type Set}
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
  |> list.try_map(fn(char) {
    case char {
      ">" -> Ok(Right)
      "^" -> Ok(Up)
      "<" -> Ok(Left)
      "v" -> Ok(Down)
      _ -> Error(UnknownCharError(char))
    }
  })
}

pub type MoveMemory {
  MoveMemory(cp: Coordinates, visited_nodes: Set(Coordinates))
}

fn move_to_coordinates(acc: MoveMemory, direction: Direction) -> MoveMemory {
  let new_position = case direction {
    Up -> Coordinates(x: acc.cp.x, y: acc.cp.y + 1)
    Down -> Coordinates(x: acc.cp.x, y: acc.cp.y - 1)
    Left -> Coordinates(x: acc.cp.x - 1, y: acc.cp.y)
    Right -> Coordinates(x: acc.cp.x + 1, y: acc.cp.y)
  }

  MoveMemory(
    cp: new_position,
    visited_nodes: set.insert(acc.visited_nodes, new_position),
  )
}

pub fn calculate_move_list(input: String) -> Result(MoveMemory, ParsingError) {
  use directions <- result.try(transform_to_direction(input))

  let initial_state =
    MoveMemory(
      cp: Coordinates(x: 0, y: 0),
      visited_nodes: set.from_list([Coordinates(x: 0, y: 0)]),
    )

  Ok(list.fold(directions, initial_state, move_to_coordinates))
}

pub fn split_directions(
  directions: List(Direction),
) -> #(List(Direction), List(Direction)) {
  list.fold(directions, #([], [], True), fn(acc, dir) {
    let #(first, second, should_add_to_first) = acc
    case should_add_to_first {
      True -> #([dir, ..first], second, False)
      False -> #(first, [dir, ..second], True)
    }
  })
  |> fn(acc) {
    let #(first, second, _) = acc
    #(list.reverse(first), list.reverse(second))
  }
}

pub fn part1(input: String) -> Solution {
  case calculate_move_list(input) {
    Ok(move_list) -> Answer(set.size(move_list.visited_nodes))
    Error(_) -> Answer(0)
    // Or handle error appropriately
  }
}

pub fn part2(input: String) -> Solution {
  case transform_to_direction(input) {
    Ok(direction_list) -> {
      let #(santas_moves, robot_moves) = split_directions(direction_list)

      let initial_state =
        MoveMemory(
          cp: Coordinates(x: 0, y: 0),
          visited_nodes: set.from_list([Coordinates(x: 0, y: 0)]),
        )

      let santas_move_list =
        list.fold(santas_moves, initial_state, move_to_coordinates)

      let robot_move_list =
        list.fold(robot_moves, initial_state, move_to_coordinates)

      let all_visited =
        set.union(santas_move_list.visited_nodes, robot_move_list.visited_nodes)

      Answer(set.size(all_visited))
    }
    Error(_) -> Answer(0)
  }
}
