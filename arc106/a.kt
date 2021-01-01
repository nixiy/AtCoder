import java.io.PrintWriter
import kotlin.math.pow

val pw = PrintWriter(System.out)

fun main(args: Array<String>) {
    val N = readLong()
    var flag = false

    run loop@ {
        for (i in 1..37) {
            for (j in 1..25) {
                val ans = (3.0.toBigDecimal().pow(i) + 5.0.toBigDecimal().pow(j)).toLong()
                if (ans == N) {
                    println("$i $j")
                    flag = true
                    return@loop
                }
            }
        }
    }

    if (!flag) {
        println(-1)
    }

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
fun readSplitDouble() = readLine()!!.split(" ").map{ it.toDouble() }
fun readSplitLong() = readLine()!!.split(" ").map{ it.toLong() }
fun readSplitUInt() = readLine()!!.split(" ").map{ it.toUInt() }
fun readSplitInt() = readLine()!!.split(" ").map{ it.toInt()}
fun readSplitStr() = readLine()!!.split(" ")
fun readLong() = readLine()!!.toLong()
fun readInt() = readLine()!!.toInt()
fun dust() = readLine()

fun println(value: Any, total: Int) {
    pw.println(value)
}
