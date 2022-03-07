fun main(args: Array<String>) {
    val (d, t, s) = readLine()!!.split(" ").map(String::toInt)

    if (d/s.toDouble() <= t.toDouble()) {
        println("Yes")
    } else {
        println("No")
    }
}
