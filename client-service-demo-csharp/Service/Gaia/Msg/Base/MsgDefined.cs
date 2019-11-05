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
using System.Collections.Generic;
using System.Reflection;

namespace GameBox.Service
{
    /// <summary>
    /// 消息定义
    /// </summary>
    internal abstract class MsgDefined : IMsgDefined
    {
        /// <summary>
        /// 消息实际类型字典
        /// </summary>
        private readonly Dictionary<string, Type> MsgTypesByName;

        /// <summary>
        /// 允许根据消息Id获取消息信息的映射
        /// </summary>
        private readonly Dictionary<uint, IMsgInfo> MsgInfosByID;

        /// <summary>
        /// 允许根据消息名字获取消息信息的映射
        /// </summary>
        private readonly Dictionary<string, IMsgInfo> MsgInfosByName;

        /// <summary>
        /// 构建一个新的消息定义实例
        /// </summary>
        protected MsgDefined()
        {
            MsgTypesByName = new Dictionary<string, Type>();
            MsgInfosByID = new Dictionary<uint, IMsgInfo>();
            MsgInfosByName = new Dictionary<string, IMsgInfo>();
        }

        /// <summary>
        /// 根据消息名字获取消息详细信息
        /// </summary>
        /// <param name="msgName">规定的消息名字</param>
        /// <returns>消息信息</returns>
        public IMsgInfo GetMsgByName(string msgName)
        {
            IMsgInfo info;
            return !MsgInfosByName.TryGetValue(msgName, out info) ? null : info;
        }

        /// <summary>
        /// 根据全名获取消息信息
        /// </summary>
        /// <param name="msgName">规定的消息全名</param>
        /// <returns>消息信息</returns>
        public IMsgInfo GetMsgByFullName(string msgName)
        {
            var temp = msgName.Split('.');
            return temp.Length > 0 ? GetMsgByName(temp[temp.Length - 1]) : null;
        }

        /// <summary>
        /// 根据消息Id获取消息信息
        /// </summary>
        /// <param name="msgId">消息Id</param>
        /// <returns>消息信息</returns>
        public IMsgInfo GetMsgById(uint msgId)
        {
            IMsgInfo info;
            return !MsgInfosByID.TryGetValue(msgId, out info) ? null : info;
        }

        /// <summary>
        /// 增加消息定义
        /// </summary>
        /// <param name="msgId">消息Id</param>
        /// <param name="msgName">消息名字</param>
        public void AddMsgDefined(ushort msgId, string msgName, Type type)
        {
            //if (msgId >= 65536)
            //{
            //    throw new ArgumentOutOfRangeException("msgId", "Param [msgId] is larger than 65535");
            //}

            //Type type = null;

            //if (MsgTypesByName.ContainsKey(msgName))
            //{
            //    type = MsgTypesByName[msgName];
            //}
            var msgType = GetMsgType(type);

            var msgInfo = DecorateMsgInfo(new MsgInfo
            {
                Id = msgId,
                Name = msgName,
                MsgType = msgType,
                ProtocolType = type
            });

            MsgInfosByID[msgId] = msgInfo;
            MsgInfosByName[msgName] = msgInfo;
        }

        public void RemoveMsgDefined(uint msgId)
        {
            IMsgInfo msginfo;
            if (MsgInfosByID.TryGetValue(msgId, out msginfo))
            {
                MsgInfosByID.Remove(msgId);
                MsgInfosByName.Remove(msginfo.Name);
            }
        }

        /// <summary>
        /// 获取消息类型
        /// </summary>
        /// <param name="type">消息的真实类型</param>
        /// <returns>消息类型</returns>
        protected virtual MsgTypes GetMsgType(Type type)
        {
            throw new NotSupportedException();
        }

        /// <summary>
        /// 对MsgInfo进行修饰处理
        /// </summary>
        /// <param name="msgInfo">消息信息</param>
        /// <returns>修饰后的消息信息</returns>
        protected virtual MsgInfo DecorateMsgInfo(MsgInfo msgInfo)
        {
            return msgInfo;
        }

        /// <summary>
        /// 从CurrentDomain中加载符合条件的消息类型
        /// </summary>
        /// <param name="predicate">条件筛选器</param>
        protected void LoadMsgTypesWithCurrentDomain(Predicate<Type> predicate)
        {
            if (predicate == null)
            {
                throw new ArgumentException("Param [predicate] cannot be null", "predicate");
            }

            // Only search current DLL
            var assembly = Assembly.GetExecutingAssembly();
            foreach (var type in assembly.GetTypes())
            {
                if (predicate(type))
                {
                    MsgTypesByName.Add(type.Name, type);
                }
            }
        }
    }
}
