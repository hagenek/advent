import java.io.File

fun moveToInt(move: String): Int {
    return when (move) {
        "A" ->  1
        "B" ->  2
        "C" ->  3
        "X" ->  1
        "Y" ->  2
        "Z" ->  3
        else -> 0
    }
}
// Rock = 1, Paper = 2, Scissor = 3
// Player Two's perspective
// Rock vs. Scissor = -2 Loss
// Rock vs. Paper = -1 Win
// Rock vs. Rock = 0 Draw
// Paper vs. Paper = 0 Draw
// Paper vs. Rock = 1 Loss
// Paper vs. Scissor = -1 Win
// Scissor vs. Rock = 2 Win
// Scissor vs. Paper = 1 Loss
// Scissor vs. Scissor = 0 Draw

enum class Outcome {
    DRAW,
    WIN,
    LOSS
}

fun chooseWinningMove(rpsPlayerOne: String, outComeWanted: String): Int {

    when (rpsPlayerOne) {
        "A" -> {
            return when (outComeWanted) {
                "X" -> 3 + 0
                "Y" -> 1 + 3
                "Z" -> 2 + 6
                else -> 0
            }
        }
        "B" -> {
            return when (outComeWanted) {
                "X" -> 1 + 0
                "Y" -> 2 + 3
                "Z" -> 3 + 6
                else -> 0
            }
        }
        else -> {
            return when (outComeWanted) {
                "X" -> 2 + 0
                "Y" -> 3 + 3
                "Z" -> 1 + 6
                else -> 0
            }
        }
    }
}

fun calculateScore(rpsPlayerOne: String, rpsPlayerTwo: String): Int {
    val pOne = moveToInt(rpsPlayerOne)
    val pTwo = moveToInt(rpsPlayerTwo)

    return when {
        pOne == pTwo -> 3 + pTwo
        (pOne - pTwo) == -1 || pOne - pTwo == 2 -> 6 + pTwo
        else -> 0 + pTwo
    }
}


fun dayTwo() {
    val realDataList = mutableListOf<String>()
    File("resources/day_two.txt").forEachLine { realDataList.add(it) }
    realDataList.map { Pair(it.substring(0, 1), it.substring(2,3)) }

    val demoDataList = mutableListOf<String>()
    File("resources/day_two_demo.txt").forEachLine { demoDataList.add(it) }
    demoDataList.map { Pair(it.substring(0, 1), it.substring(2,3)) }

    val scoreList = realDataList
            .map { Pair(it.substring(0, 1), it.substring(2,3)) }
            .map { calculateScore(it.first, it.second) }
    println("Part one: ${scoreList.sum()}")

    val scoreListPartTwo = realDataList
            .map { Pair(it.substring(0, 1), it.substring(2,3)) }
            .map { chooseWinningMove(it.first, it.second) }
    println("Part two sum: ${scoreListPartTwo.sum()})
}

fun main(args: Array<String>) {
    println("Results of part one day two:")
    dayTwo()
    // Try adding program arguments via Run/Debug configuration.
    // Learn more about running applications: https://www.jetbrains.com/help/idea/running-applications.html.
    println("Program arguments: ${args.joinToString()}")
}