pub struct Solution {}

impl Solution {
    pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
        let mut vec = vec![1; nums.len()];

        for i in 1..nums.len() {
            vec[i] = vec[i - 1] * nums[i - 1];
        }

        let mut right = 1;
        for i in (0..nums.len()).rev() {
            vec[i] *= right;
            right *= nums[i];
        }

        vec
    }
}