# For example:

# (()) and ()() both result in floor 0.
# ((( and (()(()( both result in floor 3.
# ))((((( also results in floor 3.
# ()) and ))( both result in floor -1 (the first basement level).
# ))) and )())()) both result in floor -3.

defmodule ElixAdvent.Day01 do
  def zero_example do
    "(())"
  end

  def check_char(char) do
    case char do
      "(" -> 1
      ")" -> -1
      _ -> 0
    end
  end
end
