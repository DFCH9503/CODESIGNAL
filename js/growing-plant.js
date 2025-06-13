function growingPlant(upSpeed, downSpeed, desiredHeight) {
    let grownSpeed = upSpeed - downSpeed
    return Math.max(0, Math.ceil((desiredHeight - upSpeed) / grownSpeed)) + 1
}

let upSpeed = 100, downSpeed = 10, desiredHeight = 910
console.log(`The plant will reach the desired height ${desiredHeight} in`, growingPlant(upSpeed, downSpeed, desiredHeight), `day/s`)