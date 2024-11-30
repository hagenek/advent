defmodule FileUtils do
  @moduledoc """
  Solutions for Advent of Code challenges
  """

  def read_input(day) do
    File.read!("input/day#{day}.txt")
  end

  def read_input_lines(day) do
    read_input(day)
    |> String.split("\n", trim: true)
  end

  def read_input_numbers(day) do
    read_input_lines(day)
    |> Enum.map(&String.to_integer/1)
  end
end
