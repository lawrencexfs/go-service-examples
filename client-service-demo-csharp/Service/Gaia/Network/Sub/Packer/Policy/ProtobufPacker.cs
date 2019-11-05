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
    /// Protobuf包处理器
    /// </summary>
    internal static class ProtobufPacker
    {
        /// <summary>
        /// 打包
        /// </summary>
        /// <param name="packet">需要发送的数据包</param>
        /// <param name="next">投递到下一个处理器</param>
        /// <returns>打包的数据</returns>
        public static ArraySegment<byte> Packed(SentPackage packet, Func<SentPackage, ArraySegment<byte>> next)
        {
            if (packet.MsgInfo == null || packet.MsgInfo.MsgType != MsgTypes.Protobuf)
            {
                return next(packet);
            }

            var stream = new System.IO.MemoryStream(packet.SentBuffer.Array, packet.SentBuffer.Offset,
                packet.SentBuffer.Count);
           
            ProtoBuf.Meta.RuntimeTypeModel.Default.Serialize(stream, packet.Packet);
            return new ArraySegment<byte>(packet.SentBuffer.Array, packet.SentBuffer.Offset,
                (int)stream.Position);
        }

        /// <summary>
        /// 解包
        /// </summary>
        /// <param name="packet">接受数据包</param>
        /// <param name="next">投递到下一个处理器</param>
        /// <returns>解包的对象</returns>
        public static object Unpacked(ReceivePackage packet, Func<ReceivePackage, object> next)
        {
            if (packet.MsgInfo == null || packet.MsgInfo.MsgType != MsgTypes.Protobuf)
            {
                return next(packet);
            }

            // 如果这里引发了异常，那么这个协议会被抛弃
            // 所以我们不捕获这里的异常
            var stream = new System.IO.MemoryStream(packet.Body.Array, packet.Body.Offset, packet.Body.Count);
            stream.SetLength(packet.Body.Count);
            return ProtoBuf.Meta.RuntimeTypeModel.Default.Deserialize(stream, null, packet.MsgInfo.ProtocolType);
        }
    }
}
