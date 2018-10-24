'use strict'
/*
Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
*/

function sortNumber(a,b) {
    return a - b;
}

function checkPairsForSum(data,k) {
    if (data.length <= 1) return false;
    let s = data.sort(sortNumber);
    let l =0;
    let r = data.length-1;
    while (l < r) {
        let sum = s[l] + s[r];
        if(sum > k ) {
            r--;
        }else if(sum < k) {
            l++;
        }else {
            return true;
        }
    }
    return false;
}

console.log(checkPairsForSum([10,15,3,7],17));
console.log(checkPairsForSum([10,15,3,7],16));
console.log(checkPairsForSum([10,15,3,7],10));
console.log(checkPairsForSum([10,15,3,7],10));
console.log(checkPairsForSum([10,15,3,7],26));
console.log(checkPairsForSum([10,15,3,7],25));
