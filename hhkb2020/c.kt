import java.io.PrintWriter

val pw = PrintWriter(System.out)
val N = 200010

fun main(args: Array<String>) {
    dust()
    val inputs = readSplitInt()
    val disList = BooleanArray(N)
    var pastIndex = 0 //走査したindexを記憶する事で続きから探索を開始する

    inputs.forEach {
        disList[it] = true //既出箇所を1にする

        //既出箇所のフラグをON
        while(disList[pastIndex] == true) pastIndex++

        println(pastIndex)
    }
    pw.flush()
}

//便利系
fun readSplitInt() = readLine()!!.split(" ").map{ it.toInt()}
fun readSplitStr() = readLine()!!.split(" ")
fun readInt() = readLine()!!.toInt()
fun dust() = readLine()

fun println(value: Any) {
    pw.println(value)
}
