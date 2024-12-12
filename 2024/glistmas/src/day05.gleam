import gleam/dict
import gleam/int
import gleam/list
import gleam/option.{None, Some}
import gleam/string
import stropti
import types.{Answer}
import utils

fn build_rule_dict(rules: List(String)) -> dict.Dict(Int, List(Int)) {
  rules
  |> list.fold(dict.new(), fn(dict, rule) {
    let assert [number, after_number] =
      string.split(rule, "|") |> list.map(utils.int)

    dict.upsert(dict, number, fn(maybe) {
      case maybe {
        Some(existing) -> [after_number, ..existing]
        None -> [after_number]
      }
    })
  })
}

fn parse_lines(lines: List(String)) -> List(List(Int)) {
  lines
  |> list.map(fn(line) {
    line
    |> string.split(",")
    |> list.map(utils.int)
  })
}

pub fn part1(input: String) {
  let assert [rules, lines] = extract_lines_from_data(input)
  let rule_dict = build_rule_dict(rules)
  let num_lines = parse_lines(lines)

  let results =
    num_lines
    |> list.map(fn(line) {
      case is_valid_sequence(line, rule_dict) {
        True -> get_middle_element(line)
        False -> 0
      }
    })
    |> list.filter(fn(result) { result > 0 })
    |> int.sum

  Answer(results)
}

pub fn part2(input: String) {
  let assert [rules, lines] = extract_lines_from_data(input)
  let rule_dict = build_rule_dict(rules)
  let num_lines = parse_lines(lines)

  let results =
    num_lines
    |> list.filter(fn(line) { !is_valid_sequence(line, rule_dict) })
    |> list.map(fn(broken_line) {
      let valid_ordering = build_valid_ordering(broken_line, rule_dict, [])
      get_middle_element(valid_ordering)
    })
    |> int.sum

  Answer(results)
}

fn must_come_before(a: Int, b: Int, rules: dict.Dict(Int, List(Int))) -> Bool {
  case dict.get(rules, a) {
    Ok(nexts) -> list.contains(nexts, b)
    Error(_) -> False
  }
}

fn can_come_next(
  num: Int,
  remaining: List(Int),
  rules: dict.Dict(Int, List(Int)),
) -> Bool {
  list.all(remaining, fn(other) { !must_come_before(other, num, rules) })
}

fn build_valid_ordering(
  remaining: List(Int),
  rules: dict.Dict(Int, List(Int)),
  current: List(Int),
) -> List(Int) {
  case remaining {
    [] -> current
    _ -> {
      case
        list.find(remaining, fn(num) { can_come_next(num, remaining, rules) })
      {
        Ok(next) -> {
          let new_remaining = list.filter(remaining, fn(x) { x != next })
          build_valid_ordering(
            new_remaining,
            rules,
            list.append(current, [next]),
          )
        }
        Error(_) -> current
      }
    }
  }
}

fn extract_lines_from_data(data: String) {
  data
  |> string.trim
  |> string.split("\n\n")
  |> list.map(stropti.split_on_newline)
}

fn get_middle_element(l: List(Int)) {
  let middle = list.length(l) / 2
  list.index_fold(l, 0, fn(acc, element, index) {
    case index == middle {
      True -> element
      False -> acc
    }
  })
}

fn is_valid_sequence(line: List(Int), rules: dict.Dict(Int, List(Int))) -> Bool {
  case line {
    [] -> True
    [_] -> True
    [a, b, ..rest] -> {
      case dict.get(rules, a) {
        Ok(possible_nexts) -> {
          case list.contains(possible_nexts, b) {
            True -> is_valid_sequence([b, ..rest], rules)
            False -> False
          }
        }
        Error(_) -> False
      }
    }
  }
}
