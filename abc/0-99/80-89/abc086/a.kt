fun main(args: Array<String>) {
    var inpusStr = readLine()!!.split(" ").map(String::toInt)

    println(if ((inpusStr.get(0) * inpusStr.get(1)) % 2 == 0) "Even" else "Odd")
}
