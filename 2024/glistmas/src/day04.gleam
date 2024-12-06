import gleam/dict
import gleam/int
import gleam/io
import gleam/list
import gleam/string
import types.{type Solution, Answer}

pub type Coordinate {
  Coord(x: Int, y: Int)
}

pub type CoordinateMap {
  CoordinateMap(coord: Coordinate, char: String)
}

fn extract_lines_from_data(data: String) -> List(String) {
  data
  |> string.trim
  |> string.split("\n")
  |> list.map(string.trim)
  |> list.filter(fn(s) { !string.is_empty(s) })
}

fn find_xmas(count: Int, line: List(String)) -> Int {
  case line {
    ["X", "M", "A", "S", ..rest] -> find_xmas(count + 1, rest)
    [_, ..rest] -> find_xmas(count, rest)
    [] -> count
  }
}

fn find_mas(masses: List(List(Coordinate)), line: List(CoordinateMap)) {
  case line |> list.reverse {
    [
      CoordinateMap(c1, "M"),
      CoordinateMap(c2, "A"),
      CoordinateMap(c3, "S"),
      ..rest
    ] -> find_mas([[c1, c2, c3], ..masses], rest)

    [
      CoordinateMap(c1, "S"),
      CoordinateMap(c2, "A"),
      CoordinateMap(c3, "M"),
      ..rest
    ] -> find_mas([[c1, c2, c3], ..masses], rest)

    [_, ..rest] -> find_mas(masses, rest)
    [] -> masses |> list.flatten
  }
}

fn extract_coordinates(lines: List(String), x: Int) -> List(CoordinateMap) {
  lines
  |> list.index_map(fn(char, index) { CoordinateMap(Coord(x, index), char) })
}

// ex [#(0,1), #(0,2),
//  #(1, 1), #(1, 2)]

fn coordinates_to_possible_lines(
  coords: List(CoordinateMap),
  list_height: Int,
  list_width: Int,
) -> List(List(CoordinateMap)) {
  let horizontal =
    list.range(0, list_height)
    |> list.map(fn(num: Int) {
      list.filter(coords, fn(cm) { cm.coord.x == num })
    })

  let vertical =
    list.range(0, list_width)
    |> list.map(fn(num: Int) {
      list.filter(coords, fn(cm) { cm.coord.y == num })
    })

  let diagonals_down =
    coords
    |> list.group(fn(cm) { cm.coord.y - cm.coord.x })
    |> dict.values

  let diagonals_up =
    coords
    |> list.group(fn(cm) { cm.coord.y + cm.coord.x })
    |> dict.values

  let diagonals = list.append(diagonals_down, diagonals_up)

  let possible_lines_x = list.map(horizontal, coordinate_map_to_chars)
  let possible_lines_y = list.map(vertical, coordinate_map_to_chars)
  let possible_lines_xy = list.map(diagonals, coordinate_map_to_chars)

  let countx = possible_lines_x |> list.map(find_xmas(0, _)) |> int.sum
  let county = possible_lines_y |> list.map(find_xmas(0, _)) |> int.sum
  let countxy = possible_lines_xy |> list.map(find_xmas(0, _)) |> int.sum

  io.debug(#("DIAG", countxy, "VERT", county, "HORI", countx))

  list.flatten([diagonals, vertical, horizontal])
}

fn coordinate_map_to_chars(line: List(CoordinateMap)) -> List(String) {
  let line_regular = list.map(line, fn(x: CoordinateMap) { x.char })
  let line_backwards = line_regular |> list.reverse
  io.debug(#("Regular:", line_regular, "\nBackwards", line_backwards))
  let l = list.append(line_regular, line_backwards)
  l
}

fn get_all_lines(input: String) {
  let lines = extract_lines_from_data(input)
  let mx = lines |> list.map(string.to_graphemes)
  let height = list.length(mx)
  let assert [first, ..] = mx
  let width = list.length(first)
  let coordinates =
    mx
    |> list.index_map(fn(line, index) { extract_coordinates(line, index) })
    |> list.flatten

  let possible_lines_coordinate_map =
    coordinates_to_possible_lines(coordinates, height, width)

  let possible_lines =
    list.map(possible_lines_coordinate_map, coordinate_map_to_chars)

  possible_lines
}

pub fn part1(input: String) -> Solution {
  let possible_lines = get_all_lines(input)
  let count = possible_lines |> list.map(find_xmas(0, _)) |> int.sum

  Answer(count)
}

pub fn part2(_input: String) {
  panic as "Not implemented"
}
