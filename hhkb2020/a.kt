fun main(args: Array<String>) {
    val str = readLine()

    val str2 = readLine()!!
    if (str.equals("Y")) {
        println(str2.toUpperCase())
    } else {
        println(str2)
    }
}
