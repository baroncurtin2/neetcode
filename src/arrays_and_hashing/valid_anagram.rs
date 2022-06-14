use std::collections::HashMap;

pub struct Solution {}

impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        if s.len() != t.len() {
            return false;
        }

        let mut map = HashMap::new();

        for (s_char, t_char) in s.chars().zip(t.chars()) {
            *map.entry(s_char).or_insert(0) += 1;
            *map.entry(t_char).or_insert(0) -= 1;
        }

        map.values().all(|&x| x == 0)
    }
}