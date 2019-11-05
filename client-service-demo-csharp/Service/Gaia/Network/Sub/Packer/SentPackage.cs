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
    /// 发送数据包
    /// </summary>
    internal struct SentPackage
    {
        /// <summary>
        /// 消息信息
        /// </summary>
        public IMsgInfo MsgInfo;

        /// <summary>
        /// 需要打包的数据包
        /// </summary>
        public object Packet;

        /// <summary>
        /// 发送缓冲区
        /// </summary>
        public ArraySegment<byte> SentBuffer;
    }
}
