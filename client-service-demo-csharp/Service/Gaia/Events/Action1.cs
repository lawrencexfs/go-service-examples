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

namespace GameBox.Service
{
    /// <summary>
    /// 带一个参数的事件
    /// </summary>
    /// <typeparam name="T"></typeparam>
    internal class Action1<T> : ActionBase
    {
        /// <summary>
        /// 回调
        /// </summary>
        private Action<T> action;

        /// <summary>
        /// 构造函数
        /// </summary>
        /// <param name="action"></param>
        public Action1(Action<T> action)
        {
            this.action = action;
        }

        /// <summary>
        /// 响应事件
        /// </summary>
        /// <param name="objects"></param>
        public override void Fire(params object[] objects)
        {
            if (objects != null && objects.Length == 1 && objects[0] is T)
            {
                action((T)objects[0]);
            }
            else
            {
                ParamError();
            }
        }
    }
}
