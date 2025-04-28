function isIPv4Address(inputString){
    let regex = /^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$/
    return regex.test(inputString)
}

let inputString = "172.16.254.1"
console.log(isIPv4Address(inputString))