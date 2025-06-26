function electionWinners(votes, k){
    let maxVotesBefore = Math.max(...votes)
    let possibleWinners = 0

    
    if(k==0){
        for(vote of votes){
            if(vote == maxVotesBefore)
                possibleWinners++

        }
        return possibleWinners == 1 ? 1 : 0
    }
    
    for(vote of votes){
        if(maxVotesBefore < vote + k){
            possibleWinners++
        }
    }
    
    return possibleWinners
}

let votes = [2, 3, 5, 2], k =3

console.log(`With the votes of each candidate represented by [${votes}] and a remaining of ${k} vote/s, there could be`,electionWinners(votes, k) ,`winner/s`)