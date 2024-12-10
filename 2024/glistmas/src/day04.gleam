import gleam/dict.{type Dict}
import gleam/int
import gleam/list
import gleam/result
import gleam/set.{type Set}
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

fn extract_coordinates(lines: List(String), x: Int) -> List(CoordinateMap) {
  lines
  |> list.index_map(fn(char, index) { CoordinateMap(Coord(x, index), char) })
}

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

  list.flatten([diagonals, vertical, horizontal])
}

fn coordinate_map_to_chars(line: List(CoordinateMap)) -> List(String) {
  let line_regular = list.map(line, fn(x: CoordinateMap) { x.char })
  let line_backwards = line_regular |> list.reverse
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

pub type Point {
  Point(x: Int, y: Int)
}

pub fn point_add(p1: Point, p2: Point) -> Point {
  Point(x: p1.x + p2.x, y: p1.y + p2.y)
}

pub fn get_value(grid: Dict(Point, String), point: Point) -> String {
  dict.get(grid, point)
  |> result.unwrap("*")
}

fn list_of_list_to_fields(xs: List(List(String))) {
  xs
  |> list.index_map(fn(line, y) {
    line
    |> list.index_map(fn(char, x) { #(Point(x, y), char) })
  })
}

pub type Cross {
  Cross(Point, Point, Point, Point, Point)
}

fn check_point(
  acc: Set(Cross),
  point_tuple: #(Point, String),
  grid: Dict(Point, String),
) -> Set(Cross) {
  let #(curr_point, current_letter) = point_tuple

  let upper_right_point = point_add(curr_point, Point(x: 0, y: 2))
  let upper_right_letter = get_value(grid, upper_right_point)

  let middle_point = point_add(curr_point, Point(x: 1, y: 1))
  let middle_letter = get_value(grid, middle_point)

  let bottom_left_point = point_add(curr_point, Point(x: 2, y: 0))
  let bottom_left_letter = get_value(grid, bottom_left_point)

  let bottom_right_point = point_add(curr_point, Point(x: 2, y: 2))
  let bottom_right_letter = get_value(grid, bottom_right_point)

  let cross =
    Cross(
      curr_point,
      upper_right_point,
      middle_point,
      bottom_left_point,
      bottom_right_point,
    )

  case
    current_letter,
    upper_right_letter,
    middle_letter,
    bottom_left_letter,
    bottom_right_letter
  {
    "S", "M", "A", "S", "M" -> {
      set.insert(acc, cross)
    }
    "S", "S", "A", "M", "M" -> {
      set.insert(acc, cross)
    }
    "M", "S", "A", "M", "S" -> {
      set.insert(acc, cross)
    }
    "M", "M", "A", "S", "S" -> {
      set.insert(acc, cross)
    }
    _, _, _, _, _ -> acc
  }
}

pub fn part1(input: String) -> Solution {
  let possible_lines = get_all_lines(input)
  let count = possible_lines |> list.map(find_xmas(0, _)) |> int.sum

  Answer(count)
}

pub fn part2(input: String) -> Solution {
  let lines = extract_lines_from_data(input)
  let mx = lines |> list.map(string.to_graphemes)
  let points_to_check =
    list_of_list_to_fields(mx) |> list.flatten |> list.unique
  let grid = points_to_check |> dict.from_list

  let result =
    list.fold(over: points_to_check, from: set.new(), with: fn(acc, point) {
      check_point(acc, point, grid)
    })

  set.size(result) |> Answer
}
