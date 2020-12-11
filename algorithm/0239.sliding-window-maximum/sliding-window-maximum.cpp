#include <iostream>
#include <vector>
#include <set>
using namespace std;

class Solution {
public:
    vector<int> maxSlidingWindow(vector<int>& nums, int k) {
        multiset<int> window;
        vector<int> res;
        for(int i = 0; i < k; i++) {
            window.insert(nums[i]);
        }
        for(int i = k; i < nums.size(); i++) {
            res.push_back(*window.rbegin());
            window.erase(window.equal_range(nums[i-k]).first);
            window.insert(nums[i]);
        }
        res.push_back(*window.rbegin());
        return res;
    }
};