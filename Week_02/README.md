学习笔记

1. : https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
两个数组求交集，map记录小的数组，val是数量，大的循环添加到return中
优化方法：
两个排序，双指针往后判等则加入结果集
3. 滑动窗口暴力求解：https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/submissions/
老师课中讲的双向队列golang需要自己实现，没有封装好的；所以采用保留求解；
双指针法指向窗口两端，循环求最大值；
优化：
  记录最大值，滑动时最大值与出去的和新增的做对比，减少循环更新次数;

4. 删除最外层的括号: https://leetcode-cn.com/problems/remove-outermost-parentheses/
这题比较简单，初步想用栈存括号，清到最后是最外层；
用数组实现也行，最后计数器也可以；
优化到100% 要用到一些go的特性，如果字符串+=的话会不停复制占用内存，所以用slice

5. 各位相加： https://leetcode-cn.com/problems/add-digits/
暴力求解，虽然代码行数少，但总觉得循环层数会较多；
看了下数学定义，这个方式叫数根。

6. 中序遍历： https://leetcode-cn.com/problems/binary-tree-inorder-traversal/

