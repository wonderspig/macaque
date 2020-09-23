/**
  题目:
    根据一棵树的前序遍历与中序遍历构造二叉树。
    注意:
    你可以假设树中没有重复的元素。

    例如，给出

    前序遍历 preorder = [3,9,20,15,7]
    中序遍历 inorder = [9,3,15,20,7]
    返回如下的二叉树：

        3
       / \
      9  20
        /  \
       15   7 
  思路:
    递归,找中点
   效果:
    执行用时：8 ms, 在所有 Go 提交中击败了56.92%的用户
    内存消耗：4.2 MB, 在所有 Go 提交中击败了16.88%的用户
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
        if len(preorder) == 0{
            return nil
        }
        r, i := preorder[0], 0
        for ; i < len(inorder);i++{
            if inorder[i] == r{
                break
            }
        }
        left := buildTree(preorder[1:i+1], inorder[:i])
        right := buildTree(preorder[i+1:], inorder[i+1:])
        return &TreeNode{Val:r,Left:left,Right:right}
}
