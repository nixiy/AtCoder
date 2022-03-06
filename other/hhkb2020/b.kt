import kotlin.system.measureTimeMillis
 
fun main(args: Array<String>) {
    val (tate, yoko) = readLine()!!.split(" ").map(String::toInt)
    var patternNum = 0
 
    val masuList = List(tate) {
        readLine()!!
    }
 
    //横について
    for (i in 0..tate-1) {
        for (j in 0..yoko-2) {
            if (masuList[i][j].equals('.') && masuList[i][j+1].equals('.')) {
                patternNum++
            }
        }
    }
 
    //縦について
    for (i in 0..yoko-1) {
        for (j in 0..tate-2) {
            if (masuList[j][i].equals('.') && masuList[j+1][i].equals('.')) {
                patternNum++
            }
        }
    }
 
    println(patternNum)
}
