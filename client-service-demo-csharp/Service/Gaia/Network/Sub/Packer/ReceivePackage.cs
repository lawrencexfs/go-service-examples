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
    /// 接收包对象
    /// </summary>
    internal struct ReceivePackage
    {
        /// <summary>
        /// 消息信息
        /// </summary>
        public IMsgInfo MsgInfo;

        /// <summary>
        /// 需要反序列化的的数据片(包体)
        /// </summary>
        public ArraySegment<byte> Body;
    }
}
