import java.io.PrintWriter
import kotlin.math.pow
 
val pw = PrintWriter(System.out)
 
fun main(args: Array<String>) {
    val N = readSplitInt()
 
    pw.println(N.min())
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
