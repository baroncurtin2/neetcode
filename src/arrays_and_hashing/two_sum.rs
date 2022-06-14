use std::collections::HashMap;

pub struct Solution {}

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut map = HashMap::new();

        for (i, &num) in nums.iter().enumerate() {
            let num_target = target - num;

            if map.contains_key(&(num_target)) {
                return vec![map[&(num_target)], i as i32];
            }
            map.insert(num, i as i32);
        }
        vec![]
    }
}