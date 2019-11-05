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
    /// 打包策略
    /// </summary>
    internal interface IPackerPolicy
    {
        /// <summary>
        /// 打包
        /// </summary>
        /// <param name="buffer">缓冲数据流</param>
        /// <param name="msgInfo">消息信息</param>
        /// <param name="packet">输入的协议对象</param>
        void Packed(ref ArraySegment<byte> buffer, IMsgInfo msgInfo, object packet);

        /// <summary>
        /// 解包
        /// </summary>
        /// <param name="msgInfo">消息信息</param>
        /// <param name="buffer">需要解包的字节流</param>
        /// <returns>协议对象</returns>
        object Unpacked(IMsgInfo msgInfo, ArraySegment<byte> buffer);
    }
}
