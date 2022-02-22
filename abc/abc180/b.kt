import java.io.PrintWriter

val pw = PrintWriter(System.out)

fun main(args: Array<String>) {
    val N = readInt()
    val strs = readSplitLong()
    var absMax: Long = 0
    var absTotal: Long = 0
    var absPowerTotal: Long = 0


    for (i in strs) {
        val absNum = Math.abs(i)
        absTotal += absNum
        absPowerTotal += (absNum*absNum)
        absMax = Math.max(absMax, absNum)
    }

    //マンハッタン(絶対値の総加算
    println(absTotal)
    //ユークリッド(絶対値の総合加算をルート
    println(String.format("%.15f", Math.sqrt(absPowerTotal.toDouble())))
    //チェビシェフ(全ての絶対値の最大)
    println(absMax)
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
fun readInt() = readLine()!!.toInt()
fun dust() = readLine()

fun println(value: Any) {
    pw.println(value)
}
