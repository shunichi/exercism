use std::collections::HashSet;

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    let mut hash_set: HashSet<&'a str> = HashSet::new();
    let lowercase_word = word.to_lowercase();
    let sorted_word = sort_word(&lowercase_word);

    for s in possible_anagrams {
        let lowercase = s.to_lowercase();
        if lowercase_word != lowercase && sorted_word == sort_word(&lowercase) {
            hash_set.insert(s);
        }
    }
    hash_set
}

fn sort_word(word: &str) -> String {
    let mut char_vec: Vec<_> = word.chars().collect();
    char_vec.sort_unstable();
    char_vec.iter().collect()
}
