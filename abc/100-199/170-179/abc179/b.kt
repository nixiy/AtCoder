fun main(args: Array<String>) {
    val input = readLine()!!.toInt()
    var equalsCount = 0;

    for (i in 1..input) {
        val (a, b) = readLine()!!.split(" ").map(String::toInt)
        if (a == b) equalsCount++
        else equalsCount = 0
        if (equalsCount >= 3) break
    }

    println(if (equalsCount >= 3) "Yes" else "No")
}
