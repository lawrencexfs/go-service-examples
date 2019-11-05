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

namespace GameBox.Service
{
    /// <summary>
    /// 消息类型
    /// </summary>
    internal enum MsgTypes
    {
        /// <summary>
        /// 二进制消息
        /// </summary>
        Binary = 0,

        /// <summary>
        /// Protobuf消息
        /// </summary>
        Protobuf = 1,

        /// <summary>
        /// 真实数据
        /// </summary>
        Raw = 2,
    }
}
