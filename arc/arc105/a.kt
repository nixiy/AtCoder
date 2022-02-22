import java.io.PrintWriter

val pw = PrintWriter(System.out)

fun main(args: Array<String>) {
    val inputNums = readSplitInt()
    val okList = IntArray(inputNums.count())
    var isOk = false

    var total = 0
    inputNums.forEach {
        total += it
    }

//    println(total)

    //1個だけ選択する場合
    for (i in inputNums.indices) {
        if (inputNums[i] == total-inputNums[i]) isOk = true
    }

    //2個選択する場合
    if (!isOk) {
        //6通り試す
        val tes1 = inputNums[0] + inputNums[1]
        val tes2 = inputNums[0] + inputNums[2]
        val tes3 = inputNums[0] + inputNums[3]

        val tes4 = inputNums[1] + inputNums[2]
        val tes5 = inputNums[1] + inputNums[3]

        val tes6 = inputNums[2] + inputNums[3]

        val list = listOf(tes1, tes2, tes3, tes4, tes5, tes6)

        list.forEach {
            if (it == total-it) isOk = true
        }
    }

    println(if (isOk) "Yes" else "No")

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
