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
    /// 回调基类
    /// </summary>
    internal abstract class ActionBase
    {
        /// <summary>
        /// 响应事件
        /// </summary>
        /// <param name="objects">事件参数</param>
        public abstract void Fire(params object[] args);

        protected void ParamError()
        {
            throw new Exception("Param error");
        }
    }
}
