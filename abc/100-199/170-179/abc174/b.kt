fun main(args: Array<String>) {
    val (inputCount, D) = readLine()!!.split(" ").map(String::toInt)
    var okCount = 0

    repeat(inputCount) {
        val (x, y) = readLine()!!.split(" ").map(String::toDouble)
        if (Math.sqrt((x*x + y * y).toDouble()) <= D) {
            okCount++
        }
    }
    println(okCount)
}
