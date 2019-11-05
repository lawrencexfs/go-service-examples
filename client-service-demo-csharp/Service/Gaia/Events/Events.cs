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
    /// 事件系统
    /// </summary>
    internal class Events
    {
        /// <summary>
        /// 事件信息列表
        /// </summary>
        private Dictionary<string, ActionBase> actionInfos = new Dictionary<string, ActionBase>();

        /// <summary>
        /// 添加事件
        /// </summary>
        /// <typeparam name="T">事件响应类型</typeparam>
        /// <param name="key">事件名字</param>
        /// <param name="handler">事件处理方法</param>
        internal void Add<T>(string key, Action<T> handler)
        {
            if (actionInfos.ContainsKey(key))
            {
                throw new ArgumentException("Already contain this key" + key);
            }
            var actionInfo = new Action1<T>(handler);
            actionInfos.Add(key, actionInfo);
        }

        internal void Remove(string key)
        {
            if (actionInfos.ContainsKey(key))
            {
                actionInfos.Remove(key);
            }
        }

        /// <summary>
        /// 触发事件
        /// </summary>
        /// <param name="key">事件名字</param>
        /// <param name="args">事件参数</param>
        internal void Fire(string key, params object[] args)
        {
            ActionBase action;
            if (!actionInfos.TryGetValue(key, out action))
            {
                throw new Exception("Haven't find this key" + key);
            }
            action.Fire(args);
        }
    }
}
