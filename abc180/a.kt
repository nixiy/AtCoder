import java.io.PrintWriter

val pw = PrintWriter(System.out)

fun main(args: Array<String>) {
    val strs = readSplitInt()

    println(strs[0]-strs[1]+strs[2])
    pw.flush()
}

//最初公倍数を求める
fun gcd(a: Int, b: Int) :Int {
    if (b == 0) return a
    return gcd(b, a%b)
}

fun reprace(numList: MutableList<Int>, before: Int, after: Int) {
    for (i in numList.indices) {
        if (numList[i] == before) {
            numList.set(i, after)
        }
    }
}

//便利系
fun readSplitInt() = readLine()!!.split(" ").map{ it.toInt()}
fun readSplitStr() = readLine()!!.split(" ")
fun readInt() = readLine()!!.toInt()
fun dust() = readLine()

fun println(value: Any) {
    pw.println(value)
}
