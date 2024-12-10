-module(list_variants).
-export([remove_one/1]).

%% Direct binary pattern matching and built-in list operations
%% Optimized for performance over readability
remove_one([]) -> [];
remove_one([_]) -> [[]];
remove_one(L) ->
    Len = length(L),
    remove_one(L, Len, 0, [], <<>>).

remove_one(_, Len, Len, Acc, _) -> Acc;
remove_one(L, Len, I, Acc, _) ->
    %% erlang:split_binary is typically faster than lists operations for larger lists
    {Before, [_|After]} = lists:split(I, L),
    remove_one(L, Len, I + 1, [Before ++ After|Acc], <<>>).

%% Example usage from Gleam:
%%
%% pub fn remove_variants(list: List(Int)) -> List(List(Int)) {
%%   use variants <- external(list_variants, remove_one)
%%   variants
%% }
