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
    /// 消息信息接口
    /// </summary>
    internal interface IMsgInfo
    {
        /// <summary>
        /// 消息Id
        /// </summary>
        uint Id { get; }

        /// <summary>
        /// 消息名字
        /// </summary>
        string Name { get; }

        /// <summary>
        /// 消息基础类型
        /// </summary>
        MsgTypes MsgType { get; }

        /// <summary>
        /// 消息实际类型
        /// </summary>
        Type ProtocolType { get; }

        /// <summary>
        /// 消息生成器
        /// </summary>
        MethodInfo Generate { get; }
    }
}
