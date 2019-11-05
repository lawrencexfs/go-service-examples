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
    /// 盖亚网络系统基础消息定义
    /// </summary>
    internal sealed class GaiaMsgDefined : MsgDefined
    {
        /// <summary>
        /// 构建一个新的盖亚网络系统基础消息定义
        /// </summary>
        public GaiaMsgDefined()
        {
            LoadMsgType();
        }

        /// <summary>
        /// 加载消息类型
        /// </summary>
        private void LoadMsgType()
        {
            //LoadMsgTypesWithCurrentDomain((type) => typeof(IMsg).IsAssignableFrom(type) ||
                                                    //(type.FullName != null && type.FullName.StartsWith("pb.")));
        }

        /// <summary>
        /// 获取消息类型
        /// </summary>
        /// <param name="type">消息的真实类型</param>
        /// <returns>消息类型</returns>
        protected override MsgTypes GetMsgType(Type type)
        {
            if (type == null)
            {
                return MsgTypes.Raw;
            }

            //if (typeof(IMsg).IsAssignableFrom(type))
            //{
            //    return MsgTypes.Binary;
            //}

            if (type.FullName != null && type.FullName.StartsWith("pb."))
            {
                return MsgTypes.Protobuf;
            }

            if (type.IsValueType || type.IsClass)
            {
                return MsgTypes.Binary;
            }

            throw new Exception("Wrong message name , couldn't find it's prototype");
        }

        /// <summary>
        /// 对消息信息进行修饰
        /// </summary>
        /// <param name="msgInfo">消息信息</param>
        /// <returns>消息信息</returns>
        protected override MsgInfo DecorateMsgInfo(MsgInfo msgInfo)
        {
            msgInfo = base.DecorateMsgInfo(msgInfo);
            if (msgInfo.MsgType == MsgTypes.Binary && msgInfo.ProtocolType != null && msgInfo.ProtocolType != null)
            {
                msgInfo.Generate = msgInfo.ProtocolType.BaseType.GetMethod("Generate", BindingFlags.Static | BindingFlags.Public);
            }
            return msgInfo;
        }
    }
}
