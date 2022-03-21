// https://leetcode.com/problems/binary-tree-inorder-traversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */



func inorderTraversal(root *TreeNode) []int {
    
    if root == nil {
        return nil
    }
    
    if root.Left == nil && root.Right == nil {
        return []int{root.Val}
    }
    
    ret := inorderTraversal(root.Left)
    ret = append(ret , root.Val)
    ret = append(ret , inorderTraversal(root.Right)...)
    
    return ret
    
}