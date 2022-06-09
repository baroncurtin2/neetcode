use std::collections::HashMap;

pub struct Solution {}

impl Solution {
    pub fn contains_duplicate(nums: Vec<i32>) -> bool {
        let mut map = HashMap::new();
        for num in nums {
            if map.contains_key(&num) {
                return true;
            }
            map.insert(num, true);
        }
        false
    }
}