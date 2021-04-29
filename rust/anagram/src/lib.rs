use std::collections::HashSet;

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    let mut hash_set: HashSet<&'a str> = HashSet::new();
    let sorted_word = sort_word(word);

    for s in possible_anagrams {
        if word != *s && sorted_word == sort_word(s) {
            hash_set.insert(s);
        }
    }
    hash_set
}

fn sort_word(word: &str) -> String {
    let mut char_vec: Vec<_> = word.to_lowercase().chars().collect();
    char_vec.sort_unstable();
    char_vec.iter().collect()
}
