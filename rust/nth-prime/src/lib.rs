use std::convert::TryFrom;

pub fn nth(n: u32) -> u32 {
    let mut primes: Vec<u32> = vec![];
    let mut num = 2;
    let n_usize = usize::try_from(n).unwrap();
    while primes.len() <= n_usize {
        if not_divisible(num, &primes) {
            primes.push(num);
        }
        num += 1;
    }
    *primes.last().unwrap()
}

fn not_divisible(n: u32, primes: &Vec<u32>) -> bool {
    primes.iter().all(|p| n % p != 0)
}
