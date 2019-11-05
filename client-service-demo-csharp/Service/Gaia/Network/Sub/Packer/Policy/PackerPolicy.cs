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
    internal class PackerPolicy : IPackerPolicy
    {
        /// <summary>
        /// 打包策略链
        /// </summary>
        private readonly FilterChain<SentPackage, ArraySegment<byte>> packedPolicy;

        /// <summary>
        /// 解包策略链
        /// </summary>
        private readonly FilterChain<ReceivePackage, object> unpackedPolicy;

        /// <summary>
        /// 构建一个打包策略组
        /// </summary>
        public PackerPolicy()
        {
            packedPolicy = new FilterChain<SentPackage, ArraySegment<byte>>();
            unpackedPolicy = new FilterChain<ReceivePackage, object>();
        }

        /// <summary>
        /// 增加打包策略
        /// </summary>
        /// <param name="policy">打包策略</param>
        /// <returns>当前实例</returns>
        public PackerPolicy AddPackedPolicy(Func<SentPackage, Func<SentPackage, ArraySegment<byte>>, ArraySegment<byte>> policy)
        {
            packedPolicy.Add(policy);
            return this;
        }

        /// <summary>
        /// 增加解包策略
        /// </summary>
        /// <param name="policy">解包策略</param>
        /// <returns>当前实例</returns>
        public PackerPolicy AddUnpackedPolicy(Func<ReceivePackage, Func<ReceivePackage, object>, object> policy)
        {
            unpackedPolicy.Add(policy);
            return this;
        }

        /// <summary>
        /// 打包
        /// </summary>
        /// <param name="buffer">缓冲数据流</param>
        /// <param name="msgInfo">消息信息</param>
        /// <param name="packet">输入的协议对象</param>
        public void Packed(ref ArraySegment<byte> buffer, IMsgInfo msgInfo, object packet)
        {
            var sent = packedPolicy.Do(new SentPackage
            {
                MsgInfo = msgInfo,
                Packet = packet,
                SentBuffer = buffer
            }, package =>
            {
                throw GetNotSupportedException(package.MsgInfo);
            });

            buffer = sent;
        }

        /// <summary>
        /// 解包
        /// </summary>
        /// <param name="msgInfo">消息信息</param>
        /// <param name="buffer">需要解包的字节流</param>
        /// <returns>协议对象</returns>
        public object Unpacked(IMsgInfo msgInfo, ArraySegment<byte> buffer)
        {
            return unpackedPolicy.Do(new ReceivePackage
            {
                MsgInfo = msgInfo,
                Body = buffer
            }, package =>
            {
                throw GetNotSupportedException(package.MsgInfo);
            });
        }

        /// <summary>
        /// 获取不支持异常
        /// </summary>
        /// <param name="info">异常信息</param>
        /// <param name="msgId">消息Id</param>
        /// <returns>异常</returns>
        private Exception GetNotSupportedException(IMsgInfo info)
        {
            if (info != null)
            {
                throw new NotSupportedException("No suitable protocol processor for [" + info.Name +
                                                "]");
            }

            throw new NotSupportedException("No suitable protocol processor , sent model or unknow msgId");
        }
    }
}
