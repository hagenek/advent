import gleam/int
import gleam/list
import gleam/result
import gleam/string
import types.{type Solution}

pub type Expression {
  Multiply(x: Int, y: Int)
}

fn extract_line_from_data(data: String) -> String {
  data
  |> string.trim
}

fn parse_expression_part_2(
  acc: #(List(Expression), Bool),
  pieces: List(String),
) -> List(Expression) {
  let #(expressions, enabled) = acc
  case pieces, enabled {
    ["d", "o", "n", "'", "t", "(", ")", ..rest], _ ->
      parse_expression_part_2(#(expressions, False), rest)

    ["d", "o", "(", ")", ..rest], _ ->
      parse_expression_part_2(#(expressions, True), rest)

    [_, ..rest], False -> parse_expression_part_2(acc, rest)
    ["m", "u", "l", "(", x, ",", y, ")", ..rest], True ->
      parse_expression_part_2(
        #(
          [
            Multiply(
              int.parse(x) |> result.unwrap(1),
              int.parse(y) |> result.unwrap(1),
            ),
            ..expressions
          ],
          enabled,
        ),
        rest,
      )
    ["m", "u", "l", "(", x1, x2, ",", y, ")", ..rest], True ->
      parse_expression_part_2(
        #(
          [
            Multiply(
              int.parse(x1 <> x2) |> result.unwrap(1),
              int.parse(y) |> result.unwrap(1),
            ),
            ..expressions
          ],
          enabled,
        ),
        rest,
      )
    ["m", "u", "l", "(", x1, x2, x3, ",", y, ")", ..rest], True ->
      parse_expression_part_2(
        #(
          [
            Multiply(
              int.parse(x1 <> x2 <> x3) |> result.unwrap(1),
              int.parse(y) |> result.unwrap(1),
            ),
            ..expressions
          ],
          enabled,
        ),
        rest,
      )

    ["m", "u", "l", "(", x, ",", y1, y2, ")", ..rest], True ->
      parse_expression_part_2(
        #(
          [
            Multiply(
              int.parse(x) |> result.unwrap(1),
              int.parse(y1 <> y2) |> result.unwrap(1),
            ),
            ..expressions
          ],
          enabled,
        ),
        rest,
      )

    ["m", "u", "l", "(", x, ",", y1, y2, y3, ")", ..rest], True ->
      parse_expression_part_2(
        #(
          [
            Multiply(
              int.parse(x) |> result.unwrap(1),
              int.parse(y1 <> y2 <> y3) |> result.unwrap(1),
            ),
            ..expressions
          ],
          enabled,
        ),
        rest,
      )

    ["m", "u", "l", "(", x1, x2, ",", y1, y2, ")", ..rest], True ->
      parse_expression_part_2(
        #(
          [
            Multiply(
              int.parse(x1 <> x2) |> result.unwrap(1),
              int.parse(y1 <> y2) |> result.unwrap(1),
            ),
            ..expressions
          ],
          enabled,
        ),
        rest,
      )

    ["m", "u", "l", "(", x1, x2, ",", y1, y2, y3, ")", ..rest], True ->
      parse_expression_part_2(
        #(
          [
            Multiply(
              int.parse(x1 <> x2) |> result.unwrap(1),
              int.parse(y1 <> y2 <> y3) |> result.unwrap(1),
            ),
            ..expressions
          ],
          enabled,
        ),
        rest,
      )

    ["m", "u", "l", "(", x1, x2, x3, ",", y1, y2, ")", ..rest], True ->
      parse_expression_part_2(
        #(
          [
            Multiply(
              int.parse(x1 <> x2 <> x3) |> result.unwrap(1),
              int.parse(y1 <> y2) |> result.unwrap(1),
            ),
            ..expressions
          ],
          enabled,
        ),
        rest,
      )

    ["m", "u", "l", "(", x1, x2, x3, ",", y1, y2, y3, ")", ..rest], True ->
      parse_expression_part_2(
        #(
          [
            Multiply(
              int.parse(x1 <> x2 <> x3) |> result.unwrap(1),
              int.parse(y1 <> y2 <> y3) |> result.unwrap(1),
            ),
            ..expressions
          ],
          enabled,
        ),
        rest,
      )
    [_, ..rest], True -> parse_expression_part_2(acc, rest)
    [], _ -> expressions
  }
}

fn parse_expression_part_1(
  acc: List(Expression),
  pieces: List(String),
) -> List(Expression) {
  case pieces {
    ["m", "u", "l", "(", x, ",", y, ")", ..rest] ->
      parse_expression_part_1(
        [
          Multiply(
            int.parse(x) |> result.unwrap(1),
            int.parse(y) |> result.unwrap(1),
          ),
          ..acc
        ],
        rest,
      )
    ["m", "u", "l", "(", x1, x2, ",", y, ")", ..rest] ->
      parse_expression_part_1(
        [
          Multiply(
            int.parse(x1 <> x2) |> result.unwrap(1),
            int.parse(y) |> result.unwrap(1),
          ),
          ..acc
        ],
        rest,
      )
    ["m", "u", "l", "(", x1, x2, x3, ",", y, ")", ..rest] ->
      parse_expression_part_1(
        [
          Multiply(
            int.parse(x1 <> x2 <> x3) |> result.unwrap(1),
            int.parse(y) |> result.unwrap(1),
          ),
          ..acc
        ],
        rest,
      )

    ["m", "u", "l", "(", x, ",", y1, y2, ")", ..rest] ->
      parse_expression_part_1(
        [
          Multiply(
            int.parse(x) |> result.unwrap(1),
            int.parse(y1 <> y2) |> result.unwrap(1),
          ),
          ..acc
        ],
        rest,
      )

    ["m", "u", "l", "(", x, ",", y1, y2, y3, ")", ..rest] ->
      parse_expression_part_1(
        [
          Multiply(
            int.parse(x) |> result.unwrap(1),
            int.parse(y1 <> y2 <> y3) |> result.unwrap(1),
          ),
          ..acc
        ],
        rest,
      )

    ["m", "u", "l", "(", x1, x2, ",", y1, y2, ")", ..rest] ->
      parse_expression_part_1(
        [
          Multiply(
            int.parse(x1 <> x2) |> result.unwrap(1),
            int.parse(y1 <> y2) |> result.unwrap(1),
          ),
          ..acc
        ],
        rest,
      )

    ["m", "u", "l", "(", x1, x2, ",", y1, y2, y3, ")", ..rest] ->
      parse_expression_part_1(
        [
          Multiply(
            int.parse(x1 <> x2) |> result.unwrap(1),
            int.parse(y1 <> y2 <> y3) |> result.unwrap(1),
          ),
          ..acc
        ],
        rest,
      )

    ["m", "u", "l", "(", x1, x2, x3, ",", y1, y2, ")", ..rest] ->
      parse_expression_part_1(
        [
          Multiply(
            int.parse(x1 <> x2 <> x3) |> result.unwrap(1),
            int.parse(y1 <> y2) |> result.unwrap(1),
          ),
          ..acc
        ],
        rest,
      )

    ["m", "u", "l", "(", x1, x2, x3, ",", y1, y2, y3, ")", ..rest] ->
      parse_expression_part_1(
        [
          Multiply(
            int.parse(x1 <> x2 <> x3) |> result.unwrap(1),
            int.parse(y1 <> y2 <> y3) |> result.unwrap(1),
          ),
          ..acc
        ],
        rest,
      )
    [_, ..rest] -> parse_expression_part_1(acc, rest)
    [] -> acc
  }
}

fn eval(ex: Expression) {
  case ex {
    Multiply(a, b) -> a * b
  }
}

pub fn part1(input: String) -> Solution {
  let line = extract_line_from_data(input)

  types.Answer(
    parse_expression_part_1([], line |> string.to_graphemes)
    |> list.map(eval)
    |> int.sum,
  )
}

pub fn part2(input: String) -> Solution {
  let line = extract_line_from_data(input)

  types.Answer(
    parse_expression_part_2(#([], True), line |> string.to_graphemes)
    |> list.map(eval)
    |> int.sum,
  )
}
