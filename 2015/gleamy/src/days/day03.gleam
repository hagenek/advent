// day02.gleam
import aoc_types.{type Solution, Answer}
import gleam/bool
import gleam/list
import gleam/result
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
      ">" -> Right
      "^" -> Up
      "<" -> Left
      "v" -> Down
      _ -> panic as "Unknown char encountered"
    }
  })
  |> Ok
}

pub type MoveMemory {
  MoveMemory(current_position: Coordinates, visited_nodes: List(Coordinates))
}

fn move_to_coordinates(acc: MoveMemory, direction: Direction) -> MoveMemory {
  let new_position = case direction {
    Up -> Coordinates(x: acc.current_position.x, y: acc.current_position.y + 1)
    Down ->
      Coordinates(x: acc.current_position.x, y: acc.current_position.y - 1)
    Left ->
      Coordinates(x: acc.current_position.x - 1, y: acc.current_position.y)
    Right ->
      Coordinates(x: acc.current_position.x + 1, y: acc.current_position.y)
  }
  case acc.visited_nodes |> list.find(fn(node) { node == new_position }) {
    Ok(_) ->
      MoveMemory(
        current_position: new_position,
        visited_nodes: acc.visited_nodes,
      )
    Error(_) ->
      MoveMemory(current_position: new_position, visited_nodes: [
        new_position,
        ..acc.visited_nodes
      ])
  }
}

// Visited nodes are stored in a list
pub fn calculate_move_list(input: String) -> MoveMemory {
  let direction_list = case transform_to_direction(input) {
    Ok(directions) -> directions
    _ -> panic as "unknown char encountered"
  }
  direction_list
  |> list.fold(
    from: MoveMemory(current_position: Coordinates(x: 0, y: 0), visited_nodes: [
      Coordinates(x: 0, y: 0),
    ]),
    with: move_to_coordinates,
  )
}

// Split a list of Directions into alternating lists
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
  let move_list = calculate_move_list(input)
  Answer(list.length(move_list.visited_nodes))
}

pub fn part2(input: String) -> Solution {
  let assert Ok(direction_list) =
    input
    |> transform_to_direction

  let [santas_moves, robot_moves] =
    direction_list
    |> split_directions

  let santas_move_list =
    santas_moves
    |> list.fold(
      from: MoveMemory(
        current_position: Coordinates(x: 0, y: 0),
        visited_nodes: [Coordinates(x: 0, y: 0)],
      ),
      with: move_to_coordinates,
    )

  let robot_move_list =
    robot_moves
    |> list.fold(
      from: MoveMemory(
        current_position: Coordinates(x: 0, y: 0),
        visited_nodes: [Coordinates(x: 0, y: 0)],
      ),
      with: move_to_coordinates,
    )

  Answer(
    list.length(santas_move_list.visited_nodes)
    + list.length(robot_move_list.visited_nodes),
  )
}
// pub fn part2(input: String) -> Solution {
//   todo
// }
