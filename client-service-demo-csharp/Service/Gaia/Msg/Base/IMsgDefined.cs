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
    /// 消息定义
    /// </summary>
    internal interface IMsgDefined
    {
        /// <summary>
        /// 根据消息名字获取消息详细信息
        /// </summary>
        /// <param name="msgName">规定的消息名字</param>
        /// <returns>消息信息</returns>
        IMsgInfo GetMsgByName(string msgName);

        /// <summary>
        /// 根据全名获取消息信息
        /// </summary>
        /// <param name="msgName">规定的消息全名</param>
        /// <returns>消息信息</returns>
        IMsgInfo GetMsgByFullName(string msgName);

        /// <summary>
        /// 根据消息Id获取消息信息
        /// </summary>
        /// <param name="msgId">消息Id</param>
        /// <returns>消息信息</returns>
        IMsgInfo GetMsgById(uint msgId);

        void AddMsgDefined(ushort msgId, string msgName, Type type);

        void RemoveMsgDefined(uint msgId);
    }
}
