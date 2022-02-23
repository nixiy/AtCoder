fun main(args: Array<String>) {
    val input = readLine()!!
    println(if (input.takeLast(1).equals("s")) "${input}es" else "${input}s")
}
