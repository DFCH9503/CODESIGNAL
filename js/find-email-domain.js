function findEmailDomain(address){
    let regExpDomain=/@([A-Za-z0-9\-]+\.)+[A-Za-z]{2,}/
    
    let domainFound=address.match(regExpDomain)
    
    return domainFound[0].slice(1)

}

let address = "prettyandsimple@example.com"

console.log(`The email domain in the addres ${address} is :`, findEmailDomain(address))