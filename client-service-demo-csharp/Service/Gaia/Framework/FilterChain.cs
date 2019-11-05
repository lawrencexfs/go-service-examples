/*
* This file is part of the GameBox package.
*
* (c) Giant - MouGuangYi<mouguangyi@ztgame.com> , tanxiaoliang<tanxiaoliang@ztgame.com>
*
* For the full copyright and license information, please view the LICENSE
* file that was distributed with this source code.
*
* Document: http://192.168.150.238/GameBox/help/ 
*/

using System;
using System.Collections.Generic;

namespace GameBox.Service
{
    /// <summary>
    /// 过滤器链
    /// </summary>
    /// <typeparam name="TIn">输入参数</typeparam>
    /// <typeparam name="TResult">返回值</typeparam>
    internal sealed class FilterChain<TIn, TResult>
    {
        /// <summary>
        /// 过滤器链
        /// </summary>
        private readonly List<Func<TIn, Func<TIn, TResult>, TResult>> filterList;

        /// <summary>
        /// 过滤器列表
        /// </summary>
        public Func<TIn, Func<TIn, TResult>, TResult>[] FilterList
        {
            get
            {
                return filterList.ToArray();
            }
        }

        /// <summary>
        /// 堆栈 用于解决内部递归调用过滤器链所出现的问题
        /// </summary>
        private readonly Stack<int> stack;

        /// <summary>
        /// Then堆栈 用于解决内部递归调用过滤器链所出现的问题
        /// </summary>
        private readonly Stack<Func<TIn, TResult>> stackThen;

        /// <summary>
        /// 构建一个过滤器链
        /// </summary>
        public FilterChain()
        {
            stack = new Stack<int>();
            stackThen = new Stack<Func<TIn, TResult>>();
            filterList = new List<Func<TIn, Func<TIn, TResult>, TResult>>();
        }

        /// <summary>
        /// 增加一个过滤器
        /// </summary>
        /// <param name="filter">过滤器</param>
        /// <returns>过滤器链</returns>
        public FilterChain<TIn, TResult> Add(Func<TIn, Func<TIn, TResult>, TResult> filter)
        {
            filterList.Add(filter);
            return this;
        }

        /// <summary>
        /// 执行过滤器链
        /// </summary>
        /// <param name="inData">输入数据</param>
        /// <param name="then">当过滤器执行完成后执行的操作</param>
        public TResult Do(TIn inData, Func<TIn, TResult> then = null)
        {
            if (filterList.Count <= 0)
            {
                return then != null ? then.Invoke(inData) : default(TResult);
            }

            stackThen.Push(then);
            stack.Push(0);
            TResult result;
            try
            {
                result = filterList[0].Invoke(inData, Next);
            }
            finally
            {
                stack.Pop();
                stackThen.Pop();
            }
            return result;
        }

        /// <summary>
        /// 下一层过滤器链
        /// </summary>
        /// <param name="inData">输入数据</param>
        /// <returns>执行过滤器</returns>
        private TResult Next(TIn inData)
        {
            var index = stack.Pop();
            stack.Push(++index);
            var then = stackThen.Peek();
            if (index >= filterList.Count)
            {
                return then != null ? then.Invoke(inData) : default(TResult);
            }
            return filterList[index].Invoke(inData, Next);
        }
    }
}
