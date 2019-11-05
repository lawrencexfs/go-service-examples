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
    /// 二进制包处理器
    /// </summary>
    internal static class BinaryPacker
    {
        /// <summary>
        /// 发送数据流
        /// </summary>
        //[ThreadStatic]
        //private static MMStream sentStream;

        /// <summary>
        /// 接收数据流
        /// </summary>
        //[ThreadStatic]
        //private static MMStream receiveStream;

        /// <summary>
        /// 发送数据流
        /// </summary>
        //private static MMStream SentStream
        //{
        //    get { return sentStream ?? (sentStream = new MMStream()); }
        //}

        /// <summary>
        /// 接收数据流
        /// </summary>
        //private static MMStream ReceiveStream
        //{
        //    get { return receiveStream ?? (receiveStream = new MMStream()); }
        //}

        /// <summary>
        /// 打包
        /// </summary>
        /// <param name="packet">需要发送的数据包</param>
        /// <param name="next">投递到下一个处理器</param>
        /// <returns>打包的数据</returns>
        public static ArraySegment<byte> Packed(SentPackage packet, Func<SentPackage, ArraySegment<byte>> next)
        {
            if (packet.MsgInfo == null || packet.MsgInfo.MsgType != MsgTypes.Binary)
            {
                return next(packet);
            }

            byte[] buff = null;
            MsgHelper.PackParam(packet.Packet, out buff);

            Array.Copy(buff, 0, packet.SentBuffer.Array, packet.SentBuffer.Offset, buff.Length);
            //SentStream.Reset(packet.SentBuffer.Array, packet.SentBuffer.Offset, packet.SentBuffer.Count);
            //((IMsg)packet.Packet).Marshal(SentStream);
            return new ArraySegment<byte>(packet.SentBuffer.Array, packet.SentBuffer.Offset, buff.Length);
        }

        /// <summary>
        /// 解包
        /// </summary>
        /// <param name="packet">接受数据包</param>
        /// <param name="next">投递到下一个处理器</param>
        /// <returns>解包的对象</returns>
        public static object Unpacked(ReceivePackage packet, Func<ReceivePackage, object> next)
        {
            if (packet.MsgInfo == null || packet.MsgInfo.MsgType != MsgTypes.Binary)
            {
                return next(packet);
            }

            //ReceiveStream.Reset(packet.Body.Array, packet.Body.Offset, packet.Body.Count);

            var buff = new byte[packet.Body.Count];
            Array.Copy(packet.Body.Array, packet.Body.Offset, buff, 0, packet.Body.Count);

            var msg = MsgHelper.UnpackParam(packet.MsgInfo.ProtocolType, buff); // Generate(packet.MsgInfo) as IMsg;

            //if (msg == null)
            //{
            //    throw new NotSupportedException("Can not generate protocol [" + packet.MsgInfo.Name + "]");
            //}

            //msg.UnMarshal(ReceiveStream);
            return msg;
        }

        /// <summary>
        /// 对象生成
        /// </summary>
        /// <param name="msgInfo">基础消息信息</param>
        /// <returns></returns>
        private static object Generate(IMsgInfo msgInfo)
        {
            if (msgInfo.Generate != null)
            {
                return msgInfo.Generate.Invoke(null, null);
            }

            if (msgInfo.ProtocolType != null)
            {
                return Activator.CreateInstance(msgInfo.ProtocolType);
            }

            return null;
        }
    }
}
