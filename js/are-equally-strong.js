function areEquallyStrong(yourLeft, yourRight, friendsLeft, friendsRight){

    let myArms=[yourLeft,yourRight]
    let friendArms=[friendsLeft,friendsRight]
    
    let maxMyArms=Math.max(...myArms)
    let minMyArms=Math.min(...myArms)
    let maxFriendArms=Math.max(...friendArms)
    let minFriendArms=Math.min(...friendArms)
    
    return maxMyArms===maxFriendArms && minMyArms===minFriendArms
}

let yourLeft = 10, yourRight = 15, friendsLeft = 15, friendsRight = 10

console.log(areEquallyStrong(yourLeft, yourRight, friendsLeft, friendsRight))