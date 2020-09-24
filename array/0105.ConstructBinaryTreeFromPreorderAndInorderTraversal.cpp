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

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
        if(preorder.size()<=0){
            return nullptr;
        }
         return reBuild(preorder,  0,preorder.size()-1,inorder, 0, preorder.size()-1);
    }
    TreeNode* reBuild(vector<int>& preorder, int ps,int pe,vector<int>& inorder,int is,int ie){
        if(is>ie){
            return nullptr;
        }
        //先序
        TreeNode*root=new TreeNode(preorder[ps]);
        //开始寻找
        int idx=is;
        for(int i=is;i<=ie;i++){
            if(inorder[i]==preorder[ps]){
                idx=i;
                break;
            }
        }
        root->left=reBuild(preorder,  ps+1,ps+idx-is,inorder, is, idx-1);
        root->right=reBuild(preorder, ps+idx-is+1,pe,inorder, idx+1, ie);
        return root;
    }
};
