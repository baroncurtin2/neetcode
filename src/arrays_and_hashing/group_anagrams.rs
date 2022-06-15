use std::collections::HashMap;

pub struct Solution {}

impl Solution {
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut map: HashMap<String, Vec<String>> = HashMap::new();

        fn get_sorted(s: &str) -> String {
            let mut chars: Vec<char> = s.chars().collect();
            chars.sort();
            chars.iter().collect()
        }

        for s in strs {
            let sorted = get_sorted(&s);
            map.entry(sorted).or_insert(Vec::new()).push(s);
        }

        map.values().map(|v| v.clone()).collect()
    }
}