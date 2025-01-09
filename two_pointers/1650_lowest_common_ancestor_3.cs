using System;
using System.Collections.Generic;

// Definiton of a binary tree node class
// public class EduTreeNode
// {
//     public int data;
//     public EduTreeNode left;
//     public EduTreeNode right;
//     public EduTreeNode parent;

//     public EduTreeNode(int value)
//     {
//         this.data = value;
//         this.left = null;
//         this.right = null;
//         this.parent = null;
//     }
// }

public class Solution
{
    public EduTreeNode LowestCommonAncestor_HashSet(EduTreeNode p, EduTreeNode q)
    {
      var visited = new HashSet<EduTreeNode>();

      while (p != q)
      {
        if (p != null && !visited.Add(p)) return p;

        if (q != null && !visited.Add(q)) return q;

        p = p is null ? null : p.parent;
        q = q is null ? null : q.parent;
      }

      return p;
    }

    public EduTreeNode LowestCommonAncestor_TwoPointers(EduTreeNode p, EduTreeNode q)
    {
        EduTreeNode first = p, second = q;

        while(first != second)
        {
            first = first.parent == null ? q : first.parent;
            second = second.parent == null ? p : second.parent;
        }

        return first;
    }
}