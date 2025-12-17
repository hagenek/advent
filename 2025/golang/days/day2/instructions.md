# Day 2: Gift Shop

## Part 1

You get inside and take the elevator to its only other stop: the gift shop. "Thank you for visiting the North Pole!" gleefully exclaims a nearby sign. You aren't sure who is even allowed to visit the North Pole, but you know you can access the lobby through here, and from there you can access the rest of the North Pole base.

As you make your way through the surprisingly extensive selection, one of the clerks recognizes you and asks for your help.

As it turns out, one of the younger Elves was playing on a gift shop computer and managed to add a whole bunch of invalid product IDs to their gift shop database! Surely, it would be no trouble for you to identify the invalid product IDs for them, right?

They've even checked most of the product ID ranges already; they only have a few product ID ranges (your puzzle input) that you'll need to check.

### Example Input

```
11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124
```
*(The ID ranges are wrapped here for legibility; in your input, they appear on a single long line.)*

### Rules

- The ranges are separated by commas (`,`)
- Each range gives its first ID and last ID separated by a dash (`-`)
- **Invalid IDs** are made only of some sequence of digits repeated twice:
  - `55` (5 twice)
  - `6464` (64 twice)
  - `123123` (123 twice)
- Numbers don't have leading zeroes (`0101` isn't an ID at all; `101` is a valid ID that you would ignore)

### Example Analysis

| Range | Invalid IDs | Count |
|-------|------------|-------|
| `11-22` | 11, 22 | 2 |
| `95-115` | 99 | 1 |
| `998-1012` | 1010 | 1 |
| `1188511880-1188511890` | 1188511885 | 1 |
| `222220-222224` | 222222 | 1 |
| `1698522-1698528` | None | 0 |
| `446443-446449` | 446446 | 1 |
| `38593856-38593862` | 38593859 | 1 |
| Other ranges | None | 0 |

**Total sum of invalid IDs:** `1227775554`

**Question:** What do you get if you add up all of the invalid IDs?

## Part 2

The clerk quickly discovers that there are still invalid IDs in the ranges in your list. Maybe the young Elf was doing other silly patterns as well?

Now, an ID is invalid if it is made only of some sequence of digits repeated **at least twice**.

### Updated Rules

- **Invalid IDs** are made of some sequence of digits repeated at least twice:
  - `12341234` (1234 two times)
  - `123123123` (123 three times)
  - `1212121212` (12 five times)
  - `1111111` (1 seven times)

### Updated Example Analysis

| Range | Invalid IDs | Count | Notes |
|-------|------------|-------|-------|
| `11-22` | 11, 22 | 2 | Same as Part 1 |
| `95-115` | 99, 111 | 2 | Now includes 111 |
| `998-1012` | 999, 1010 | 2 | Now includes 999 |
| `1188511880-1188511890` | 1188511885 | 1 | Same as Part 1 |
| `222220-222224` | 222222 | 1 | Same as Part 1 |
| `1698522-1698528` | None | 0 | Same as Part 1 |
| `446443-446449` | 446446 | 1 | Same as Part 1 |
| `38593856-38593862` | 38593859 | 1 | Same as Part 1 |
| `565653-565659` | 565656 | 1 | New invalid ID |
| `824824821-824824827` | 824824824 | 1 | New invalid ID |
| `2121212118-2121212124` | 2121212121 | 1 | New invalid ID |

**Total sum of invalid IDs:** `4174379265`

**Question:** What do you get if you add up all of the invalid IDs using these new rules?
