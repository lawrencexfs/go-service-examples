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
using System.Reflection;

namespace GameBox.Service
{
    /// <summary>
    /// 消息信息
    /// </summary>
    internal sealed class MsgInfo : IMsgInfo
    {
        /// <summary>
        /// 消息Id
        /// </summary>
        public uint Id { get; set; }

        /// <summary>
        /// 消息名字
        /// </summary>
        public string Name { get; set; }

        /// <summary>
        /// 消息基础类型
        /// </summary>
        public MsgTypes MsgType { get; set; }

        /// <summary>
        /// 消息实际类型
        /// </summary>
        public Type ProtocolType { get; set; }

        /// <summary>
        /// 类型生成器
        /// </summary>
        public MethodInfo Generate { get; set; }
    }
}
