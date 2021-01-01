import java.io.PrintWriter

val pw = PrintWriter(System.out)
var MAX = 0
var MIN = 0

fun main(args: Array<String>) {
    val inCnt = readInt()
    val nList= readSplitInt()
    var pastGcd = nList[0]

    for (i in 1..inCnt-1) {
        pastGcd = gcd(pastGcd, nList[i])
    }

    println(pastGcd)

    pw.flush()
}

fun secondMax(numList: MutableList<Int>) :Int{
    var numListTmp = numList
    numListTmp.remove(MAX)
    return numListTmp.max()!!
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
