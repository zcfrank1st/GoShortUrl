package define

import (
    "strings"
    "math"
)

var (
    number_62 []string

    dict map[int]string

    randomIndex []int
)

func init () {
    //randomIndex = rand.Perm(62)
    number_62  = []string {
        "1","2","3","4","5","6","7","8","9","0",
        "q","w","e","r","t","y","u","i","o","p",
        "a","s","d","f","g","h","j","k","l",
        "z","x","c","v","b","n","m",
        "Q","W","E","R","T","Y","U","I","O","P",
        "A","S","D","F","G","H","J","K","L",
        "Z","X","C","V","B","N","M",
    }
}

//func main() {
//    input := 1212312313123
//    value := to62(input)
//
//    fmt.Printf("input is %d ==> result is %s \n", input, value)
//    fmt.Printf("origin: %d \n", toDecimal(value))
//}

func makeRandomDict() map[int]string {
    dict = make(map[int]string)
    for i:=0 ; i<62; i++ {
        dict[i] = number_62[randomIndex[i]]
    }
    return dict
}

func To62(i int) string {
    trans := []string{}
    dict := makeRandomDict()
    for i >= 62 {
        trans = append([]string{dict[i % 62]}, trans...)
        i = i / 62
    }
    trans = append([]string{dict[i]}, trans...)
    return strings.Join(trans, "")
}

func ToDecimal(symbol string) int{
    bytes := []byte(symbol)
    length :=len(bytes)

    var sum int
    for i:=0; i<length; i++ {
        tok := string(bytes[length - i - 1])
        sum += value2key(tok) * pow62(i)
    }
    return sum
}

func value2key(value string) int {
    n := make(map[string]int)
    for k, v := range dict {
        n[v] = k
    }
    return n[value]
}

func pow62(i int) int{
    return (int)(math.Pow(62, float64(i)))
}